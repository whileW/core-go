package httpx

import (
	"fmt"
	"testing"
)

func TestPost(t *testing.T) {
	data,err := Post("http://192.168.111.1:8888/test","",nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(data)
}
