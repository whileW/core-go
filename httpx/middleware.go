package httpx

import (
	"github.com/gin-gonic/gin"
	"github.com/whileW/core-go/log"
	"github.com/whileW/core-go/utils"
	"net/http"
	"net/http/httputil"
	"runtime"
)

// panic处理
func RecoverHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			httprequest, _ := httputil.DumpRequest(c.Request, false)
			log.GetLoger().Errorw("http_panic","module","panic","err",err,"request",string(httprequest),"stack",string(buf))
			//log.Error(pnc)
			c.AbortWithStatusJSON(500,err)
		}
	}()
	c.Next()
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.GetHeader("Origin")
		c.Header("Access-Control-Allow-Origin", utils.IF(origin == "", "*", origin).(string))
		c.Header("Access-Control-Allow-Headers", c.GetHeader("Access-Control-Request-Headers"))
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}