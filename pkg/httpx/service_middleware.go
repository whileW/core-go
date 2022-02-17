package httpx

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/whileW/core-go/utils"
	"net/http"
	"runtime"
)

func CORS(c *ServiceContext)  {
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
func RecoverHandler(ctx *ServiceContext) {
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			ctx.GetLoger().Error(fmt.Sprintf("http service panic %v",err),"module","panic","stack",string(buf))
			ctx.AbortWithStatusJSON(500,err)
		}
	}()
	ctx.Next()
}
func UserTraceId(ctx *ServiceContext)  {
	user_trace_id,_ := ctx.Cookie("user_trace_id")
	if user_trace_id == "" {
		user_trace_id = uuid.New().String()
		ctx.SetCookie("user_trace_id", user_trace_id, 12*60*60, "/", "", false, true)
	}
	ctx.SetPubLoger("user_trace_id",user_trace_id)
}