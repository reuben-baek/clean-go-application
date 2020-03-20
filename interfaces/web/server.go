package web

import (
	"golang.org/x/sys/unix"
	"net/http"
	"os"
	"os/signal"
)

type server struct {
	rootRouter
	address string
}

func Server(router rootRouter, address string) *server {
	return &server{router, address}
}

func (s *server) start(done chan<- bool) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, unix.SIGTERM)

	go func() {
		httpServer := &http.Server{
			Addr:    s.address,
			Handler: s.rootRouter,
		}

		httpServer.ListenAndServe()
	}()

	<-signalChan
	done <- true
}
