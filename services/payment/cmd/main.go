package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rvinnie/bigO/services/payment/config"
	"github.com/sirupsen/logrus"
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
		logrus.Fatal("Unable to parse config")
	}

	//Creating gRPC server
	fmt.Println(cfg.GRPC.Port)
}
