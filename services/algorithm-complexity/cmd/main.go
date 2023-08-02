package main

import (
	"github.com/joho/godotenv"
	"github.com/rvinnie/bigO/services/algorithm-complexity/config"
	"github.com/rvinnie/bigO/services/algorithm-complexity/openai_manager"
	"github.com/rvinnie/bigO/services/algorithm-complexity/transport/grpc"
	"github.com/rvinnie/bigO/services/algorithm-complexity/transport/grpc/handler"
	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
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

	//Initializing config
	cfg, err := config.InitConfig(configPath)
	if err != nil {
		logrus.Fatal("Unable to parse config", err)
	}

	// Creating chat GPT client
	openAIManager := openai_manager.NewOpenAIManager(openai.NewClient(cfg.OpenAI.Key), openai.GPT3Dot5Turbo)

	// Creating handlers
	grpcHandler := handler.NewAlgorithmComplexityHandler(openAIManager)

	// Creating gRPC server
	grpcServer := grpc.NewServer(grpcHandler)
	go func() {
		if err = grpcServer.ListenAndServe(cfg.GRPC.Port); err != nil {
			logrus.Fatalf("Error occured while running algorithm-complexity (gRPC) server: %s", err.Error())
		}
	}()
	logrus.Info("Algorithm-complexity (gRPC) server is running")

	// Gracefull shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)

	<-quit
	logrus.Info("Algorithm-complexity (gRPC) server shutting down")

	grpcServer.Stop()
}
