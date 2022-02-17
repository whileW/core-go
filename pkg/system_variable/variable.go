package system_variable

import (
	"os"
	"path/filepath"
)

var(
	// 系统名称
	SystemName 				string			=		filepath.Base(os.Args[0])
	// 运行时go语言版本
	GoVersion				string
	// 运行时环境
	Env						string			= 		defaultEnv
)

const (
	Env_TESTING 		=	"testing"
	Env_DEBUG			=	"debug"
	Env_RELEASE			=	"release"
)
var(
	defaultEnv = Env_RELEASE
)



const (
	SystemEnvVariableName_SystemName = "SN"
	SystemEnvVariableName_ENV = "ENV"
)