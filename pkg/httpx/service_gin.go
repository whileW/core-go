package httpx

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// todo 待完善

type GinServiceHandler struct {
	e *gin.Engine
	g *GinRouterGroup
}
func (h *GinServiceHandler)ServeHTTP(w http.ResponseWriter,r *http.Request) {
	h.e.ServeHTTP(w,r)
}
func NewGinServiceHandler() *GinServiceHandler {
	return &GinServiceHandler{
		e: gin.New(),
	}
}
func SetGinServiceHandler(fns ...func(rg *GinServiceHandler)) func(s *Service) {
	e := NewGinServiceHandler()
	e.Use(RecoverHandler,CORS,UserTraceId)
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

type GinRouterGroup struct {
	rg 	*gin.RouterGroup
}

func (g *GinRouterGroup)Group(relativePath string, handlers ...HandlerFunc) *GinRouterGroup {
	rg := g.rg.Group(relativePath,parseHandlers(handlers...)...)
	return &GinRouterGroup{rg:rg}
}
func (g *GinRouterGroup)GET(relativePath string, handlers ...HandlerFunc)  {
	g.rg.GET(relativePath,parseHandlers(handlers...)...)
}
func parseHandlers(handlers ...HandlerFunc) []gin.HandlerFunc {
	gin_handlers := []gin.HandlerFunc{}
	for _,t := range handlers {
		gin_handlers = append(gin_handlers, func(handler HandlerFunc) gin.HandlerFunc {
			return func(c *gin.Context) {
				handler(NewServiceContext(c))
			}
		}(t))
	}
	return gin_handlers
}