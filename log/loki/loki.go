package loki

import (
	"github.com/sirupsen/logrus"
	"github.com/whileW/core-go/conf"
)

func SetLokiOutPut()  {
	addr := get_loki_setting().GetStringd("addr","http://127.0.0.1:3100")
	labels := map[string]string{
		"app":conf.GetConf().SysSetting.SystemName,
	}

	promtailHook,err := NewPromtailHook(addr, labels)
	if err != nil {
		panic(err)
	}
	logrus.AddHook(promtailHook)
}

func get_loki_setting() *conf.Settings {
	return conf.GetConf().Setting.GetChildd("loki")
}