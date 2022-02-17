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
	req_loger	*log.Loger
}

func NewServiceContext(c *gin.Context) *ServiceContext {
	loger := log.WithKV("req_id",uuid.New().String())
	return &ServiceContext{
		Context:c,
		loger: loger.Clone(),
		req_loger: loger.Clone(),
	}
}
// 获取clone得log
func (ctx *ServiceContext)GetLoger() *log.Loger {
	return ctx.loger.Clone()
}
func (ctx *ServiceContext)GetReqLoger() *log.Loger {
	return ctx.req_loger
}
func (ctx *ServiceContext)SetPubLoger(arg ...interface{}) {
	ctx.loger.WithKV(arg...)
	ctx.req_loger.WithKV(arg...)
}
func (ctx *ServiceContext)GetDB() *gorm.DB {
	return orm.GetDB().WithContext(ctx.GetLoger().WithModule("orm").ValueCtx())
}
// todo 待完善
func (ctx *ServiceContext)Ok()  {
	
}