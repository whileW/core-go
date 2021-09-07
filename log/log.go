package log

import (
	"github.com/sirupsen/logrus"
	"github.com/whileW/core-go/conf"
	"github.com/whileW/core-go/log/loki"
	"os"
)
//todo 集成loki
func init()  {
	//设置日志format
	logrus.SetFormatter(&logrus.JSONFormatter{})
	//设置输出
	//设置级别
	logrus.SetLevel(logrus.InfoLevel)

	if conf.GetConf().SysSetting.Env == "debug" {
		logrus.SetOutput(os.Stdout)
	}
	if get_log_setting().GetBoold("loki",false){
		loki.SetLokiOutPut()
	}
}

func get_log_setting() *conf.Settings {
	return conf.GetConf().Setting.GetChildd("log")
}

func GetLoger() *Loger {
	return &Loger{}
}