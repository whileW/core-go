package conf

import (
	"github.com/whileW/core-go/utils"
	"os"
	"strings"
)

func initEnv(config *Config) {
	config.SysSetting.Env = strings.TrimSpace(os.Getenv("ENV"))
	config.SysSetting.HttpAddr = strings.TrimSpace(os.Getenv("HTTPADDR"))
	config.SysSetting.RpcAddr = strings.TrimSpace(os.Getenv("RPCADDR"))
	config.SysSetting.Host = strings.TrimSpace(os.Getenv("HOST"))
	config.SysSetting.SystemName = strings.TrimSpace(os.Getenv("SYSTEMNAME"))

	config.SysSetting.ConfFrom = utils.IF(os.Getenv("CONFFROM")=="","file",os.Getenv("CONFFROM")).(string)
	config.SysSetting.SetDefault()
}