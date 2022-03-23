package httpx

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
type GinServiceHandler struct {
	e *gin.Engine
	disabled_req_body_log_path []string
}
func (h *GinServiceHandler)ServeHTTP(w http.ResponseWriter,r *http.Request) {
	// todo 记录日志
	h.e.ServeHTTP(w,r)
}
//func (h *GinServiceHandler)DisabledReqBodyLog(path ...string)  {
//	h.disabled_req_body_log_path = append(h.disabled_req_body_log_path, path...)
//}
func NewGinServiceHandler() *GinServiceHandler {
	e := gin.New()
	e.Use(setServiceContext())
	return &GinServiceHandler{
		e:e,
		disabled_req_body_log_path:[]string{},
	}
}
func SetGinServiceHandler(fns ...func(rg *GinServiceHandler)) func(s *Service) {
	e := NewGinServiceHandler()
	e.Use(Middleware_ReqLog,Middleware_RecoverHandler,Middleware_UserTraceId,Middleware_CORS)
	for _,t := range fns {
		t(e)
	}
	return SetServiceHandler(e)
}

func (h *GinServiceHandler)Use(middleware ...HandlerFunc)  {
	h.e.Use(parseHandlers(middleware...)...)
}
func (h *GinServiceHandler)Group(relativePath string, handlers ...HandlerFunc) *GinRouterGroup {
	rg := h.e.Group(relativePath,parseHandlers(handlers...)...)
	return &GinRouterGroup{rg:rg}
}
func (h *GinServiceHandler)GET(relativePath string, handlers ...HandlerFunc) {
	h.e.GET(relativePath,parseHandlers(handlers...)...)
}
func (h *GinServiceHandler)POST(relativePath string, handlers ...HandlerFunc) {
	h.e.POST(relativePath,parseHandlers(handlers...)...)
}
func (h *GinServiceHandler)Any(relativePath string, handlers ...HandlerFunc) {
	h.e.Any(relativePath,parseHandlers(handlers...)...)
}

type GinRouterGroup struct {
	rg 	*gin.RouterGroup
}
func (g *GinRouterGroup)Use(middleware ...HandlerFunc)  {
	g.rg.Use(parseHandlers(middleware...)...)
}
// todo 禁用请求日志
//func (g *GinRouterGroup)GroupDisabledReqBodyLog(relativePath string, handlers ...HandlerFunc) *GinRouterGroup {
//	rg := g.rg.Group(relativePath,parseHandlers(handlers...)...)
//	return &GinRouterGroup{rg:rg}
//}
func (g *GinRouterGroup)Group(relativePath string, handlers ...HandlerFunc) *GinRouterGroup {
	rg := g.rg.Group(relativePath,parseHandlers(handlers...)...)
	return &GinRouterGroup{rg:rg}
}
func (g *GinRouterGroup)GET(relativePath string, handlers ...HandlerFunc)  {
	g.rg.GET(relativePath,parseHandlers(handlers...)...)
}
func (g *GinRouterGroup)POST(relativePath string, handlers ...HandlerFunc)  {
	g.rg.POST(relativePath,parseHandlers(handlers...)...)
}
func (g *GinRouterGroup)Any(relativePath string, handlers ...HandlerFunc)  {
	g.rg.Any(relativePath,parseHandlers(handlers...)...)
}
func parseHandlers(handlers ...HandlerFunc) []gin.HandlerFunc {
	gin_handlers := []gin.HandlerFunc{}
	for _,t := range handlers {
		gin_handlers = append(gin_handlers, func(c *gin.Context) {
			t(getServiceContext(c))
		})
	}
	return gin_handlers
}