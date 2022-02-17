package loki

import (
	"github.com/sirupsen/logrus"
	"github.com/whileW/core-go/pkg/conf"
	"github.com/whileW/core-go/pkg/system_variable"
	"sync"
)

func Initialize() error {
	addr := conf.GetStringd("log.loki",defalutLokiAddr)
	if addr != "" {
		return initialize(addr)
	}
	return nil
}

//// 基于flag初始化
//func init_by_flag(addr string)  {
//	if err := initialize(addr);err != nil{
//		fmt.Println(fmt.Sprintf("initialize loki by flag error. addr:%s, err:%v",addr,err))
//		initonce = sync.Once{}	// 基于flag启动失败，重置initonce（后面基于conf启动）
//		return
//	}
//}
//
//func InitByConf() error {
//	addr := conf.GetStringd("log.loki","")
//	if addr != "" {
//		return initialize(addr)
//	}
//	return nil
//}


var initonce = sync.Once{}
func initialize(addr string) (err error) {
	initonce.Do(func() {
		default_labels := map[string]string{
			"app":system_variable.SystemName,
		}
		hook,e := NewPromtailHook(addr, default_labels)
		if e != nil {
			err = e
			return
		}
		logrus.AddHook(hook)
	})

	return err
}