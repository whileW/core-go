package loki

import "github.com/whileW/core-go/pkg/flag"

var defalutLokiAddr = ""

func init()  {
	flag.Register(&flag.StringFlag{
		Name:"loki",
		Usage:"--loki=127.0.0.1:3100，开启loki日志存储",
		EnvVar:"LOKI",
		Variable:&defalutLokiAddr,
	})
}
