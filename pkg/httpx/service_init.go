package httpx

import (
	"github.com/whileW/core-go/pkg/flag"
)

func init()  {
	flag.Register(&flag.IntFlag{
		Name: "http_port",
		Usage: "--http_port=80",
		EnvVar: "HTTP_PORT",
		Action: func(key string, fs *flag.FlagSet) {
			DefaultListenPort = int(fs.Int(key))
		},
	})
}
