package httpx

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/whileW/core-go/utils"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

func setServiceContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := NewServiceContext(c)
		c.Set("serviceContext",ctx)
	}
}
func getServiceContext(c *gin.Context) *ServiceContext {
	return c.MustGet("serviceContext").(*ServiceContext)
}

func Middleware_CORS(c *ServiceContext)  {
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
func Middleware_RecoverHandler(ctx *ServiceContext) {
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
func Middleware_UserTraceId(ctx *ServiceContext)  {
	user_trace_id,_ := ctx.Cookie("user_trace_id")
	if user_trace_id == "" {
		user_trace_id = uuid.New().String()
		ctx.SetCookie("user_trace_id", user_trace_id, 12*60*60, "/", "", false, true)
	}
	ctx.SetPubLoger("user_trace_id",user_trace_id)
}
const (
	//最大打印gin resp_body长度  1mb
	MAX_PRINT_GIN_RESP_BODY_LEN  	=  1048576
	//最大打印gin req_body 长度
	MAX_PRINT_GIN_REQ_BODY_LEN		=	1048576
)
func Middleware_ReqLog(ctx *ServiceContext)  {
	//loger := ctx.GetReqLoger()
	var (
		req_time = time.Now()
		client_ip = ctx.ClientIP()
		method = ctx.Request.Method
		path = ctx.Request.RequestURI
	)

	//clone req body
	if d,err := ioutil.ReadAll(ctx.Request.Body);err == nil {
		ctx.Set("req_body_log",d)
		ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(d))
	}

	ctx.Next()

	loger := ctx.GetLoger().WithModule("req_log")
	loger.WithKV("req_time",req_time,"client_ip",client_ip,"method",method,"path",path)

	if v,ok := ctx.Get("disable_req_body_log"); !ok || v.(string) != "1" {
		req_body := ctx.MustGet("req_body_log").([]byte)
		if len(req_body) <= MAX_PRINT_GIN_REQ_BODY_LEN {
			loger.WithKV("req_body",string(req_body))
		}
	}

	loger.WithKV("resp_status_code",ctx.Writer.Status())

	//处理时间
	loger.WithDuration(req_time)

	loger.Info(fmt.Sprintf("%s %s %d",method,ctx.Request.RequestURI,ctx.Writer.Status()))
}