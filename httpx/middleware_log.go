package httpx

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/whileW/core-go/log"
	"io/ioutil"
	"time"
)

const (
	//最大打印gin resp_body长度  1mb
	MAX_PRINT_GIN_RESP_BODY_LEN  	=  1048576
	//最大打印gin req_body 长度
	MAX_PRINT_GIN_REQ_BODY_LEN		=	1048576
)

type ginRespBodyLoger struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}
func (w ginRespBodyLoger) Write(b []byte) (int, error) {
	//memory copy here!
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}

//追加请求头日志
func AppendReqHeadLog(name ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		req_head_log := []string{}
		v,ok := c.Get("req_head_log")
		if ok {
			if v,ok := v.([]string);ok {
				req_head_log = v
			}
		}
		req_head_log = append(req_head_log, name...)
		c.Set("req_head_log",req_head_log)
	}
}
//禁用请求体日志
func DisableReqBodyLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("disable_req_body_log","1")
	}
}
//重写请求体内容
func RewriteReqBody(c *gin.Context,body []byte)  {
	c.Set("req_body_log",body)
}
//禁用返回体日志
func DisableRespBodyLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("disable_resp_body_log","1")
	}
}

func GetGinContextReqId(c *gin.Context) string {
	return c.MustGet("req_id").(string)
}

//gin请求日志中间件
//todo req_id
func EnableGinLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		loger := log.GetLoger()
		start := time.Now()
		req_id := uuid.New().String()
		c.Set("req_id",req_id)
		loger.WithModule("reqlog")
		loger.WithKV("req_time",start)
		loger.WithKV("client_ip",c.ClientIP())
		loger.WithKV("req_method",c.Request.Method)
		loger.WithKV("req_path",c.Request.URL.Path)
		loger.WithKV("req_para",c.Request.RequestURI)
		loger.WithReqId(c)
		//clone req body
		if d,err := ioutil.ReadAll(c.Request.Body);err == nil {
			c.Set("req_body_log",d)
			c.Request.Body = ioutil.NopCloser(bytes.NewReader(d))
		}


		gin_body_loger := ginRespBodyLoger{
			bodyBuf: bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = gin_body_loger

		c.Next()

		if v,ok := c.Get("req_head_log");ok {
			if v,ok := v.([]string);ok {
				for _,k := range v {
					loger.WithKV(k,c.Request.Header.Get(k))
				}
			}
		}

		if v,ok := c.Get("disable_req_body_log"); !ok || v.(string) != "1" {
			req_body := c.MustGet("req_body_log").([]byte)
			if len(req_body) <= MAX_PRINT_GIN_REQ_BODY_LEN {
				loger.WithKV("req_body",string(req_body))
			}
		}

		if v,ok := c.Get("disable_resp_body_log"); !ok || v.(string) != "1" {
			resp_body := gin_body_loger.bodyBuf.Bytes()
			if len(resp_body) <= MAX_PRINT_GIN_RESP_BODY_LEN {
				loger.WithKV("resp_body",string(resp_body))
			}
		}

		loger.WithKV("resp_status_code",c.Writer.Status())

		//处理时间
		loger.WithDuration(start)

		loger.Infow("reqlog")
	}
}