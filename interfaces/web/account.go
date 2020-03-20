package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/reuben-baek/clean-go-application/application"
)

type AccountRouter struct {
	app *application.AccountApplication
}

func NewAccountRouter(app *application.AccountApplication) *AccountRouter {
	h := &AccountRouter{app}
	return h
}

func (h *AccountRouter) Routes() []Route {
	return []Route{
		{"GET", "/:id", []gin.HandlerFunc{h.get}},
	}
}

func (h *AccountRouter) get(ctx *gin.Context) {
	id := ctx.Param("id")
	account, _ := h.app.Find(id)
	ctx.String(200, fmt.Sprintf("Hello, %+v", account))
}
