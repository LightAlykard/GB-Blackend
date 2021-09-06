package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/LightAlykard/GB-Blackend/app"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)

	logger.Info("Starting the application")

	// TODO get port from settings file or environment
	port := "8080"

	r := mux.NewRouter()
	r.HandleFunc("/", app.HomeHandler())

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	shutdown := make(chan error, 1)

	server := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: r,
	}

	go func() {
		err := server.ListenAndServe()
		shutdown <- err
	}()
	logger.Infof("The service is ready to listen and serve")

	select {
	case killSignal := <-interrupt:
		switch killSignal {
		case os.Interrupt:
			logger.Info("Got SIGINT...")
		case syscall.SIGTERM:
			logger.Info("Got SIGTERM...")
		case syscall.SIGINT:
			logger.Info("Got SIGINT...")
		}
	case <-shutdown:
		logger.Info("Got an error...")
	}

	logger.Infof("The service is stopping...")
	err := server.Shutdown(context.Background())
	if err != nil {
		logger.Infof("Got an error during service shutdown: %v", err)
	}
}
