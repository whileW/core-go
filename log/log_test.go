package log

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogLevel(t *testing.T)  {
	//logrus.Panic("panic")
	//logrus.Fatal("fatal")
	logrus.Error("error")
	logrus.Warn("warn")
	logrus.Info("info")
	logrus.Debug("debug")
	logrus.Trace("trace")
}

func TestLogPrint(t *testing.T)  {
	GetLoger().Infow("test","info","信息",1,"test")
}