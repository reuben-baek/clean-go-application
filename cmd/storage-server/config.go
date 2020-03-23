package main

import (
	"github.com/gin-gonic/gin"
	"github.com/reuben-baek/clean-go-application/application"
	"github.com/reuben-baek/clean-go-application/infrastructure/inmemory"
	"github.com/reuben-baek/clean-go-application/interfaces/web"
)

func init() {
	accountRepository := inmemory.NewAccountRepository()
	accountApp := application.NewAccountApplication(accountRepository)
	accountRouter := web.NewAccountRouter(accountApp)
	web.RootRouter.Handle(accountRouter)

	engine := gin.Default()
	web.RootRouter.Init(engine)
}
