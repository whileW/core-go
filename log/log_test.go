package log

import (
	"github.com/sirupsen/logrus"
	"github.com/whileW/core-go/log/loki"
	"testing"
	"time"
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

func TestSetLokiOutPut(t *testing.T) {
	loki.SetLokiOutPut()
	//GetLoger().Errorw("test","key1","value")
	for i:=0;i<10 ;i++  {
		GetLoger().Infow("测试2","key","test","module","reqlog")
		time.Sleep(5*time.Second)
	}
}
func TestLoger_Copy(t *testing.T) {
	l := GetLoger().WithKV("key","test1")
	l2 := l.Copy()
	l.WithKV("key2","test2")
	l2.Infow("测试2")
	l.Infow("测试")
}