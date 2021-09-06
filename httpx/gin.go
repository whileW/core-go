package httpx

import (
	"github.com/gin-gonic/gin"
	"github.com/whileW/core-go/conf"
)

func NewGin() *gin.Engine {
	if conf.GetConf().SysSetting.Env != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	//开启gin
	r := gin.New()
	// 请求日志记录
	r.Use(EnableGinLog())
	// 跨域
	r.Use(Cors())
	//捕获异常
	r.Use(RecoverHandler)
	return r
}
