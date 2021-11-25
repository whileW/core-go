package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"os"
)

func get_nacos_addr() string {
	nacos_addr := os.Getenv("NACOSADDR")
	if nacos_addr == "" {
		return "nacos"
	}
	return nacos_addr
}


// 获取nacos初始化配置
func get_c_s_config() (*constant.ClientConfig,[]constant.ServerConfig) {
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
			get_nacos_addr(),
			8848,
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos"),
		),
	}
	return c,s
}

// 服务注册、发现
func GetNamingClient() (naming_client.INamingClient,error) {
	c,s := get_c_s_config()
	return clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  c,
			ServerConfigs: s,
		},
	)
}


// 统一配置
func GetConfigClient() (config_client.IConfigClient,error) {
	c,s := get_c_s_config()
	return clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  c,
			ServerConfigs: s,
		},
	)
}