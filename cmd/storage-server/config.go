package main

import (
	"github.com/gin-gonic/gin"
	"github.com/reuben-baek/clean-go-application/application"
	"github.com/reuben-baek/clean-go-application/infrastructure/inmemory"
	"github.com/reuben-baek/clean-go-application/interfaces/web"
	"github.com/reuben-baek/clean-go-application/lib/webserver"
)

func init() {
	accountRepository := inmemory.NewAccountRepository()
	accountApp := application.NewDefaultAccountApplication(accountRepository, nil)
	accountRouter := web.NewAccountRouter(accountApp)
	webserver.Handle(accountRouter)

	engine := gin.Default()
	webserver.Init(engine)
}
