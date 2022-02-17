package api

import (
	"github.com/whileW/core-go/example/api/user"
	"github.com/whileW/core-go/pkg/httpx"
)

func NewApiService() *httpx.Service {
	httpServer := httpx.NewServer(
		httpx.SetGinServiceHandler(
			user.RegisterUserHandler,
			),
	)
	httpServer.SetDefaultListenPort(8080)
	return httpServer
}

