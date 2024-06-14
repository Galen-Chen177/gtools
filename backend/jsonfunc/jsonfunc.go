package jsonfunc

import (
	"net/url"
	"os"

	"gtools-wails/backend/utils"
)

type IJsonFunc interface {
	JsonFile2Struct(string) string
	Json2Struct(string) string

	UrlParse(string) string
}

func NewJsonFunc() IJsonFunc {
	return &JsonFunc{}
}

type JsonFunc struct{}

func (j *JsonFunc) JsonFile2Struct(filename string) string {
	res, err := os.ReadFile(filename)
	if err != nil {
		return utils.Resp("", err)
	}
	return j.Json2Struct(string(res))
}

func (j *JsonFunc) Json2Struct(jsonStr string) string {
	genJson, err := New([]byte(jsonStr), "")
	if err != nil {
		return utils.Resp("", err)
	}
	bytes, err := genJson.WriteGo()
	if err != nil {
		return utils.Resp("", err)
	}
	return utils.Resp(string(bytes), nil)
}

func (j *JsonFunc) UrlParse(urlOrigin string) string {
	enEscapeUrl, err := url.QueryUnescape(urlOrigin)
	return utils.Resp(enEscapeUrl, err)
}
