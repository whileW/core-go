package handler

import (
	"fmt"
	"github.com/whileW/core-go/example/model"
	"github.com/whileW/core-go/pkg/httpx"
)

func RegisterTestHandlers(rg *httpx.GinRouterGroup)  {
	r := rg.Group("test")
	{
		r.GET("info", testInfoHandler)
	}
}

func testInfoHandler(ctx *httpx.ServiceContext)  {
	d,err := model.SearchUserFirst(ctx.GetDB(),
		model.SearchUserOption_Name("小明"),
	)
	if err != nil {
		//....
	}
	//....
	fmt.Println(d)
}