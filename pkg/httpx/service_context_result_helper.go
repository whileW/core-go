package httpx

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code 	int		        `json:"code"`
	Data 	interface{} 	`json:"data"`
	Msg  	interface{}    `json:"msg"`
}

const (
	ResCode_SUCCESS 				= 	0
	//服务器内部逻辑异常
	ResCode_ERROR   				= 	500
	//请求参数检查错误
	ResCode_ParamterError			=	501
	ResCode_NoFind					=	404
	//登陆异常
	ResCode_LoginFailure			=	401
	//鉴权失败
	ResCode_Unauthorized			=	402
)

func (ctx *ServiceContext)Result(ResCode int,ResMsg interface{},ResData interface{})  {
	ctx.JSON(http.StatusOK,Response{
		Code:ResCode,
		Msg:ResMsg,
		Data:ResData,
	})
	ctx.Abort()
}
func (ctx *ServiceContext)NoFindResult()  {
	ctx.Status(http.StatusNotFound)
	ctx.Abort()
}
func (ctx *ServiceContext)Ok(c *gin.Context)  {
	ctx.Result(ResCode_SUCCESS,"操作成功",nil)
}
func (ctx *ServiceContext)OkWithMessage(c *gin.Context,ResMsg interface{}) {
	ctx.Result(ResCode_SUCCESS, ResMsg, nil)
}
func (ctx *ServiceContext)OkWithData(c *gin.Context,ResData interface{}) {
	ctx.Result(ResCode_SUCCESS, "操作成功", ResData)
}
func (ctx *ServiceContext)OkDetailed(c *gin.Context,ResMsg interface{},ResData interface{}) {
	ctx.Result(ResCode_SUCCESS,ResMsg, ResData)
}

func (ctx *ServiceContext)Fail(c *gin.Context) {
	ctx.Result(ResCode_ERROR,"操作失败",nil)
}
func (ctx *ServiceContext)FailWithMessage(c *gin.Context,ResMsg interface{}) {
	ctx.Result(ResCode_ERROR, ResMsg,nil)
}
func (ctx *ServiceContext)FailWithData(c *gin.Context,ResData interface{}) {
	ctx.Result(ResCode_ERROR, "操作失败",ResData)
}
func (ctx *ServiceContext)FailWithDetailed(c *gin.Context,ResCode int, ResMsg interface{}, ResData interface{}) {
	ctx.Result(ResCode, ResMsg, ResData)
}