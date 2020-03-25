package main

import (
	"github.com/reuben-baek/clean-go-application/lib/webserver"
)

func main() {
	webServer := webserver.Server(webserver.RootRouter, "localhost:8080")
	webServer.Start()
}
