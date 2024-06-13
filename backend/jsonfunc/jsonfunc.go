package jsonfunc

import "os"

type IJsonFunc interface {
	JsonFile2Struct(string) (string, error)
	Json2Struct(string) (string, error)
}

func NewJsonFunc() IJsonFunc {
	return &JsonFunc{}
}

type JsonFunc struct{}

func (j *JsonFunc) JsonFile2Struct(filename string) (string, error) {
	res, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return j.Json2Struct(string(res))
}

func (j *JsonFunc) Json2Struct(jsonStr string) (string, error) {
	genJson, err := New([]byte(jsonStr), "")
	if err != nil {
		return "", err
	}
	bytes, err := genJson.WriteGo()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
