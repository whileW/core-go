package system_variable

import (
	"github.com/whileW/core-go/pkg/flag"
	"github.com/whileW/core-go/pkg/util/xdebug"
	"runtime"
)

func init() {
	// 系统名称
	flag.Register(&flag.StringFlag{
		Name:     "sn",
		Usage:    "--sn=system",
		EnvVar:   SystemEnvVariableName_SystemName,
		Action: func(key string, fs *flag.FlagSet) {
			SystemName = fs.String(key)
		},
	})
	// 运行时go语言版本
	GoVersion = runtime.Version()
	// 运行环境：release\debug\testing
	if xdebug.IsTestingMode() {
		Env = Env_TESTING
	}else {
		flag.Register(&flag.StringFlag{
			Name:  "env",
			Usage: "--env=release",
			EnvVar: SystemEnvVariableName_ENV,
			Action: func(key string, fs *flag.FlagSet) {
				v := fs.String(key)
				if v == Env_DEBUG {
					Env = Env_DEBUG
				}else {
					Env = Env_RELEASE
				}
			},
		})
	}
}