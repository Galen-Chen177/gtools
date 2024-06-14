package utils

import (
	"errors"
	"fmt"
	"net/url"
	"testing"
)

func TestResp(t *testing.T) {
	// {"data":"{\"a\":\"123\"}","err":"456"}
	fmt.Println(Resp(`{"a":"123"}`, errors.New("456")))
}

func TestUrlP(t *testing.T) {
	var urlStr string = "傻了吧:%:%@163& .html.html"
	escapeUrl := url.QueryEscape(urlStr)
	fmt.Println("编码:", escapeUrl)

	enEscapeUrl, _ := url.QueryUnescape(`https://translate.google.com/?hl=zh-CN&sl=auto&tl=en&text=%E5%8C%85%E5%90%AB&op=translate`)
	fmt.Println("解码:", enEscapeUrl)
}
