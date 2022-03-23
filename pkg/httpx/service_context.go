package httpx

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/whileW/core-go/pkg/log"
	"github.com/whileW/core-go/pkg/orm"
	"gorm.io/gorm"
)

type HandlerFunc func(*ServiceContext)
type ServiceContext struct {
	*gin.Context
	loger 	*log.Loger
	//req_loger	*log.Loger
	//is_disabled_req_body_log 	bool
}
// todo pool
func NewServiceContext(c *gin.Context) *ServiceContext {
	return &ServiceContext{
		loger:log.WithKV("req_id",uuid.New().String()),
		Context:c,
	}
}


//func (ctx *ServiceContext)disabled_req_body_log() *ServiceContext {
//	ctx.is_disabled_req_body_log = true
//	return ctx
//}
// 获取clone得log
func (ctx *ServiceContext)GetLoger() *log.Loger {
	return ctx.loger.Clone()
}
//func (ctx *ServiceContext)GetReqLoger() *log.Loger {
//	return ctx.req_loger
//}
func (ctx *ServiceContext)SetPubLoger(arg ...interface{}) {
	ctx.loger.WithKV(arg...)
	//ctx.req_loger.WithKV(arg...)
}
func (ctx *ServiceContext)GetDB() *gorm.DB {
	return orm.GetDB().WithContext(ctx.GetLoger().WithModule("orm").ValueCtx())
}