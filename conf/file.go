package conf

import (
	"fmt"
	"github.com/whileW/core-go/utils"
	"os"
	"runtime"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//自动查找配置文件--自动上面三级查找配置文件
func initFile(config *Config) {
	conf_file_name := utils.IF(os.Getenv("CFNAME") == "", "config", os.Getenv("CFNAME")).(string)
	conf_file_type := utils.IF(os.Getenv("CFTYPE") == "", "yaml", os.Getenv("CFNAME")).(string)

	v := viper.New()
	v.SetConfigName(conf_file_name)
	v.SetConfigType(conf_file_type)
	v.AddConfigPath("./")
	v.AddConfigPath("../")
	v.AddConfigPath("../../")
	v.AddConfigPath("../../../")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %v \n", err))
		fmt.Println(fmt.Sprintf("Fatal error config file: %v \n", err))
		return
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		config.AnalysisSetting(v.AllSettings())
	})
	config.AnalysisSetting(v.AllSettings())
	//通过runtime.KeepAlive 保证v不被垃圾回收
	runtime.KeepAlive(v)
}
