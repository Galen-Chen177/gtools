package utils

import (
	"errors"
	"fmt"
	"testing"
)

func TestResp(t *testing.T) {
	// {"data":"{\"a\":\"123\"}","err":"456"}
	fmt.Println(Resp(`{"a":"123"}`, errors.New("456")))
}
