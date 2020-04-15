package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/reuben-baek/clean-go-application/application"
	"github.com/reuben-baek/clean-go-application/lib/webserver"
)

type AccountRouter struct {
	app application.AccountApplication
}

func NewAccountRouter(app application.AccountApplication) *AccountRouter {
	h := &AccountRouter{app}
	return h
}

func (h *AccountRouter) Routes() []webserver.Route {
	return []webserver.Route{
		webserver.NewRoute("GET", "/:id", h.get),
		webserver.NewRoute("PUT", "/:id", h.put),
	}
}

func (h *AccountRouter) get(ctx *gin.Context) {
	id := ctx.Param("id")
	account, _ := h.app.FindOne(id)
	ctx.JSON(200, account)
}

func (h *AccountRouter) put(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.app.Save(application.NewAccount(id))
	if err != nil {
		ctx.String(500, fmt.Sprintf("Internal Error. %v", err))
	} else {
		ctx.Status(200)
	}
}
