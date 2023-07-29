package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/rvinnie/bigO/services/gateway/config"
	"github.com/rvinnie/bigO/services/gateway/transport/rest"
	"github.com/rvinnie/bigO/services/gateway/transport/rest/handler"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	configPath = "./config"
)

func main() {
	// Adding logger
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Initializing env variables
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file")
	}

	// Initializing config
	cfg, err := config.InitConfig(configPath)
	if err != nil {
		logrus.Fatal("Unable to parse config file")
	}

	handlers := handler.NewHandler()

	restServer := rest.NewServer(cfg, handlers.Init(cfg))
	go func() {
		if err = restServer.Run(); err != http.ErrServerClosed {
			logrus.Fatalf("error occured while running gateway (HTTP) server: %s", err.Error())
		}
	}()
	logrus.Info("Gateway (HTTP) server is running")

	// Gracefull shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)

	<-quit

	logrus.Info("Gateway (HTTP) server shutting down")
	if err = restServer.Stop(context.Background()); err != nil {
		logrus.Errorf("Error on gateway (HTTP) server shutting down: %s", err.Error())
	}
}
