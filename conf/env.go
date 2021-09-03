package conf

import (
	"github.com/whileW/core-go/utils"
	"os"
)

func initEnv(config *Config) {
	config.SysSetting.Env = os.Getenv("ENV")
	config.SysSetting.HttpAddr = os.Getenv("HTTPADDR")
	config.SysSetting.RpcAddr = os.Getenv("RPCADDR")
	config.SysSetting.Host = os.Getenv("HOST")

	config.SysSetting.ConfFrom = utils.IF(os.Getenv("CONFFROM")=="","file",os.Getenv("CONFFROM")).(string)
}