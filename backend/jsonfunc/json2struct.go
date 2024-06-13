package jsonfunc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"
	"unicode"
	"unicode/utf8"
)

const (
	goDefaultTpl = `type {{.Name}} {{if .IsArray}}[]{{end}}struct {
`
	goArrayTpl = `_ {{if .IsArray}}[]{{end}}struct {
`
	goATplEmpty = `_ {{if .IsArray}}[]{{end}}struct {
}
`
)

// Model ...
type Model struct {
	Writer      io.Writer
	Name        string
	Data        interface{}
	WithExample bool
	Format      bool
	Convert     bool
}

// New ...
func New(byte []byte, name string) (m *Model, err error) {
	var data interface{}
	if err = json.Unmarshal(byte, &data); err != nil {
		return
	}
	return &Model{
		Writer: os.Stdout,
		Data:   data,
		Name:   filter(name),
		Format: true,
	}, nil
}

// Get ...
func Get(url string) ([]byte, string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", err
	}
	req.Header.Add("Accept", "application/json")
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer func() {
		_ = r.Body.Close()
	}()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, "", err
	}
	return b, getName(url), err
}

// WriteGo ...
func (m *Model) WriteGo() (b []byte, err error) {
	if m.Format {
		var buf bytes.Buffer
		m.Writer = &buf
		m.print(func(ms map[string]interface{}) {
			m.parseMap(ms)
		})
		b, err = format.Source(buf.Bytes())
		org := m.Writer
		if err == nil {
			//_, _ = org.Write(b)
		} else {
			_, _ = io.Copy(org, &buf)
		}
		m.Writer = org
	} else {
		m.print(func(ms map[string]interface{}) {
			m.parseMap(ms)
		})
	}
	return
}

func (m *Model) print(convert func(map[string]interface{})) {
	structNum := 0
	defer func() {
		for i := 0; i < structNum; i++ {
			_, _ = fmt.Fprintln(m.Writer, "}")
		}
		_, _ = fmt.Fprintln(m.Writer, "}")
	}()

	switch v := m.Data.(type) {
	case []interface{}:
		printTpl(m.Writer, goDefaultTpl, m.Name, true)
	init:
		for _, node := range v {
			switch node.(type) {
			case []interface{}:
				if len(node.([]interface{})) == 0 {
					printTpl(m.Writer, goATplEmpty, m.Name, true)
				} else {
					v = node.([]interface{})
					printTpl(m.Writer, goArrayTpl, m.Name, true)
					structNum++
					goto init
				}
			default:
				if len(node.(map[string]interface{})) == 0 {
					printTpl(m.Writer, goATplEmpty, m.Name, false)
				}
				convert(node.(map[string]interface{}))
			}
		}
	case float64:
		break
	default:
		printTpl(m.Writer, goDefaultTpl, m.Name, false)
		convert(m.Data.(map[string]interface{}))
	}
}

func (m *Model) parseMap(ms map[string]interface{}) {
	keys := getSortedKeys(ms)
	for _, k := range keys {
		m.parse(ms[k], k)
	}
}

func (m *Model) parse(data interface{}, k string) {
	switch v := data.(type) {
	case string:
		if m.Convert {
			t, converted := parseType(v)
			m.printType(k, v, t, converted)
		} else {
			m.printType(k, v, "string", false)
		}
	case bool:
		m.printType(k, v, "bool", false)
	case float64:
		//json parser always returns a float for number values, check if it is an int value
		if float64(int64(v)) == v {
			m.printType(k, v, "int64", false)
		} else {
			m.printType(k, v, "float64", false)
		}
	case int64:
		m.printType(k, v, "int64", false)
	case []interface{}:
		if len(v) > 0 {
			switch vv := v[0].(type) {
			case string:
				m.printType(k, v[0], "[]string", false)
			case float64:
				//json parser always returns a float for number values, check if it is an int value
				if float64(int64(v[0].(float64))) == v[0].(float64) {
					m.printType(k, v[0], "[]int64", false)
				} else {
					m.printType(k, v[0], "[]float64", false)
				}
			case bool:
				m.printType(k, v[0], "[]bool", false)
			case []interface{}:
				m.parse(vv[0], k)
				//m.printObject(k, "[]struct", func() { m.parse(vv[0], k) })
			case map[string]interface{}:
				m.printObject(k, "[]struct", func() { m.parseMap(vv) })
			default:
				//fmt.Printf("unknown type: %T", vv)
				m.printType(k, nil, "interface{}", false)
			}
		} else {
			m.printType(k, nil, "[]interface{}", false)
		}
	case map[string]interface{}:
		m.printObject(k, "struct", func() { m.parseMap(v) })
	default:
		m.printType(k, nil, "interface{}", false)
	}
}

func (m *Model) printType(key string, value interface{}, t string, converted bool) {
	name := filter(key)
	if converted {
		key += ",string"
	}
	if m.WithExample {
		_, _ = fmt.Fprintf(m.Writer, "%s %s `json:\"%s\"` // %v\n", name, t, key, value)
	} else {
		_, _ = fmt.Fprintf(m.Writer, "%s %s `json:\"%s\"`\n", name, t, key)
	}
}

func (m *Model) printObject(n string, t string, f func()) {
	_, _ = fmt.Fprintf(m.Writer, "%s %s {\n", filter(n), t)
	f()
	_, _ = fmt.Fprintf(m.Writer, "} `json:\"%s\"`\n", n)
}

func getName(u string) string {
	p, err := url.Parse(u)
	if err != nil {
		return "Data"
	}
	s := strings.Split(p.Path, "/")
	if len(s) < 1 {
		return "Data"
	}
	return strings.Title(s[len(s)-1])
}

func printTpl(w io.Writer, tplData string, name string, isArray bool) {
	tmpl, err := template.New("test").Parse(tplData)
	if err != nil {
		panic(err)
	}

	data := struct {
		Name    string
		IsArray bool
	}{
		name,
		isArray,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func parseType(value string) (string, bool) {
	if _, err := time.Parse(time.RFC3339, value); err == nil {
		return "time.Time", false
	} else if ip := net.ParseIP(value); ip != nil {
		return "net.IP", false
	} else if _, err := strconv.ParseInt(value, 10, 64); err == nil {
		return "int64", true
	} else if _, err := strconv.ParseFloat(value, 64); err == nil {
		return "float64", true
	} else if _, err := strconv.ParseBool(value); err == nil {
		return "bool", true
	} else {
		return "string", false
	}
}

func filter(name string) string {
	if name == "" {
		name = "Data"
	}
	newString := ""
	for _, r := range name {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			newString += string(r)
		} else {
			newString += " "
		}
	}
	newString = strings.Title(newString)
	newString = strings.Replace(newString, " ", "", -1)
	newString = strings.Replace(newString, "Url", "URL", -1)
	newString = strings.Replace(newString, "Uri", "URI", -1)
	newString = strings.Replace(newString, "Id", "ID", -1)

	r, _ := utf8.DecodeRuneInString(name)
	if !unicode.IsLetter(r) && !(r == '_') {
		newString = "_" + newString
	}

	return newString
}

func getSortedKeys(m map[string]interface{}) (keys []string) {
	for key := range m {
		keys = append(keys, key)
	}
	sort.Sort(ByIDFirst(keys))
	return
}

// ByIDFirst ...
type ByIDFirst []string

func (p ByIDFirst) Len() int { return len(p) }
func (p ByIDFirst) Less(i, j int) bool {
	if p[i] == "id" {
		return true
	} else if p[j] == "id" {
		return false
	}
	return p[i] < p[j]
}
func (p ByIDFirst) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
