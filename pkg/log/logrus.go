package log

import (
	"github.com/sirupsen/logrus"
	"github.com/whileW/core-go/pkg/system_variable"
	"os"
)

func init()  {
	//设置日志format
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if system_variable.Env != system_variable.Env_RELEASE {
		logrus.SetOutput(os.Stdout)
	}
}
