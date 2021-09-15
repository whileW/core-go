package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/whileW/core-go/conf"
	"github.com/whileW/core-go/httpx"
	"net/http"
	"time"
)

// Server is http server.
type Server struct {
	router *gin.Engine
	server *http.Server
}
func New() *Server {
	s := &Server{}

	s.router = httpx.NewGin()

	s.server = &http.Server{
		Addr:    ":" + conf.GetConf().SysSetting.HttpAddr,
		Handler: s.router,
	}

	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	return s
}
// Close close the server.
// todo 待测试服务优雅退出
func (s *Server) Close() {
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(timeoutCtx); err != nil {
		fmt.Println(err)
	}
}

func main()  {
	s := New()
	s.router.GET("ok", func(c *gin.Context) {
		httpx.Ok(c)
	})
	s.router.GET("ok/with_data", func(c *gin.Context) {
		httpx.OkWithData(c,"data")
	})
	for {
		select {}
	}
}
