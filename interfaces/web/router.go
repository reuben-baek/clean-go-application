package web

import "github.com/gin-gonic/gin"

type rootRouter struct {
	*gin.Engine
}

func RootRouter(engine *gin.Engine) *rootRouter {
	return &rootRouter{Engine: engine}
}

func (g *rootRouter) Handle(r Router) {
	for _, route := range r.Routes() {
		g.Engine.Handle(route.method, route.path, route.handlers...)
	}
}

type Route struct {
	method   string
	path     string
	handlers []gin.HandlerFunc
}

type Router interface {
	Routes() []Route
}
