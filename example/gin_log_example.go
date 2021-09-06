package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/whileW/core-go/conf"
	"github.com/whileW/core-go/httpx"
	"github.com/whileW/core-go/log"
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
	server := New()
	log.SetPrometheusOutPut(server.router.Group(""))
	for {
		select {}
	}
}
