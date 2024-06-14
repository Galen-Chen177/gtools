package jsonfunc

import (
	"fmt"
	"log"
	"testing"
)

func TestJson2struct(t *testing.T) {
	test, err := New([]byte(`[[{"id":123},[{"test":false,"hello":"abc","msg":{}}]]]`), "")
	if err != nil {
		log.Fatalln(err.Error())
	}
	if bytes, err := test.WriteGo(); err != nil {
		log.Fatalln(err.Error())
	} else {
		fmt.Println(string(bytes))
	}
}

func TestJf(t *testing.T) {
	jf := NewJsonFunc()
	fmt.Println(jf.Json2Struct(`{"a":"123"}`))
}
