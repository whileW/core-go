package model

import (
	"fmt"
	"testing"
)

func TestGen(t *testing.T) {
	m := []*ModelJsonStruct{
		{
			TableName:"user",
			Fields:[]*ModelFieldJsonStruct{
				{
					Name:"name",
					Type:"string",
					Size:256,
				},
				{
					Name:"age_now",
					Type:"int",
					Comment:"现在的年龄",
				},
			},
			IsNeedUUID:true,
		},
	}
	fs,err := Gen(m)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _,t := range fs {
		fmt.Println(string(t.Data))
		fmt.Println("-----------------")
	}
}