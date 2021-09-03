package main

import (
	"fmt"
	"github.com/whileW/core-go/conf"
	"time"
)

//go run example/conf_nacos_example.go
func main()  {
	c := conf.GetConf()
	fmt.Println(c)
	fmt.Println(c.Setting.GetStringd("test",""))
	time.Sleep(time.Second*60)
}
