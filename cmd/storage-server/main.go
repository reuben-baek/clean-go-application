package main

import "github.com/reuben-baek/clean-go-application/interfaces/web"

func main() {
	webServer := web.Server(web.RootRouter, "localhost:8080")
	webServer.Start()
}
