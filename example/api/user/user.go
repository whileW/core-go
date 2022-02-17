package user

import (
	"github.com/whileW/core-go/example/api/user/handler"
	"github.com/whileW/core-go/pkg/httpx"
)

func RegisterUserHandler(sh *httpx.GinServiceHandler) {
	g := sh.Group("user")
	handler.RegisterTestHandlers(g)
}