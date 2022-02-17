package httpx

import (
	"fmt"
	"github.com/whileW/core-go/pkg/conf"
	"net/http"
	"strconv"
)

// 默认监听端口
var DefaultListenPort = 80
// 启动http服务 -- 由Application来启动
type Service struct {
	opts 	[]func(s *Service)
	ser 	*http.Server
	e 		http.Handler
}
// flag、config、log、orm 加载之前加载
func NewServer(opts ...func(s *Service)) *Service {
	s := &Service{
		opts: opts,
	}
	return s
}

// flag、config、log、orm 加载完毕后加载
func (s *Service)StartUp() error {
	for _,t := range s.opts {
		t(s)
	}
	if s.e == nil {
		panic("请配置httpHandler")
	}
	http_port := conf.GetIntd("http.port",DefaultListenPort)
	s.ser = &http.Server{
		Addr:    ":" + strconv.Itoa(http_port),
		Handler: s.e,
	}
	fmt.Println(fmt.Sprintf("StartUp HTTP Server at %d ...",http_port))
	return s.ser.ListenAndServe()
}
func (s *Service)Stop() error {
	return nil
}

// 修改默认得监听端口
// 优先级：Default < ENV < Flag < CONFIG
func (s *Service)SetDefaultListenPort(port int)  {
	DefaultListenPort = port
}
func SetServiceHandler(e http.Handler) func(s *Service) {
	return func(s *Service) {
		s.e = e
	}
}
func (s *Service)SetServiceHandler(e http.Handler)  {
	s.e = e
}