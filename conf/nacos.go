package conf

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/whileW/core-go/utils"
	"os"
	"runtime"
)

func initNacos(config *Config) {
	var (
		data_id = utils.IF(os.Getenv("NACOSDATAID") == "",config.SysSetting.SystemName,os.Getenv("NACOSDATAID")).(string)
		server_addr = os.Getenv("NACOSADDR")
		env     = utils.IF(config.SysSetting.Env != "release","debug",config.SysSetting.Env).(string)
	)
	c := get_nacos_client(server_addr)

	var changeData = func(confContent string) {
		s := map[string]interface{}{}
		if err := json.Unmarshal([]byte(confContent),&s);err != nil{
			fmt.Println(fmt.Sprintf("nacos序列化配置失败:content:%s,err:%v",confContent,err))
			return
		}
		fmt.Println("nacos config changed")
		config.AnalysisSetting(s)
	}

	if err := c.ListenConfig(vo.ConfigParam{
		DataId:data_id,
		Group:env,
		OnChange: func(namespace, group, dataId, data string) {
			changeData(data)
		},
	});err != nil{
		panic(fmt.Sprintf("监听nacos配置失败:%v",err))
	}
	//初始化获取配置
	content,err := c.GetConfig(vo.ConfigParam{
		DataId:data_id,
		Group:env,
	})
	if err != nil {
		panic(fmt.Sprintf("获取nacos配置失败:%v",err))
	}
	changeData(content)
	runtime.KeepAlive(c)
}

func get_nacos_client(addr string) config_client.IConfigClient {
	c := constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		//constant.WithLogDir("/tmp/nacos/log"),
		//constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithRotateTime("1h"),
		constant.WithMaxAge(3),
		constant.WithLogLevel("debug"),
	)
	s := []constant.ServerConfig{
		*constant.NewServerConfig(
			addr,
			8848,
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos"),
		),
	}

	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  c,
			ServerConfigs: s,
		},
	)
	if err != nil {
		panic(fmt.Sprintf("初始化nacos异常:%v",err))
	}
	return client
}