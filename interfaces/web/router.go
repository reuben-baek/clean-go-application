package web

import "github.com/gin-gonic/gin"

var RootRouter = &rootRouter{}

type rootRouter struct {
	*gin.Engine
	routers []Router
}

func newRootRouter(engine *gin.Engine, routers ...Router) *rootRouter {
	r := &rootRouter{routers: routers}
	r.Init(engine)
	return r
}

func (g *rootRouter) Init(engine *gin.Engine) {
	g.Engine = engine
	g.setRoutes()
}

func (g *rootRouter) setRoutes() {
	for _, router := range g.routers {
		for _, route := range router.Routes() {
			g.Engine.Handle(route.method, route.path, route.handlers...)
		}
	}
}

func (g *rootRouter) Handle(r Router) {
	g.routers = append(g.routers, r)
}

type route struct {
	method   string
	path     string
	handlers []gin.HandlerFunc
}

func Route(method string, path string, handlers ...gin.HandlerFunc) route {
	return route{
		method,
		path,
		handlers,
	}
}

type Router interface {
	Routes() []route
}
