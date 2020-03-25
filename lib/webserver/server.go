package webserver

import (
	"golang.org/x/sys/unix"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type server struct {
	*rootRouter
	address string
}

func Server(router *rootRouter, address string) *server {
	return &server{router, address}
}

func (s *server) Start() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, unix.SIGTERM)

	go func() {
		httpServer := &http.Server{
			Addr:    s.address,
			Handler: s.rootRouter,
		}

		log.Fatal(httpServer.ListenAndServe())
	}()

	<-signalChan
}
