package log

import (
	"github.com/sirupsen/logrus"
	"github.com/whileW/core-go/conf"
	"os"
)
//todo 集成loki
func init()  {
	var (
		format = get_log_setting().GetStringd("format","json")
	)
	//设置日志format
	switch format {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	default:
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	//设置输出
	//设置级别
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func get_log_setting() *conf.Settings {
	return conf.GetConf().Setting.GetChildd("log")
}

func GetLoger() *Loger {
	return &Loger{}
}