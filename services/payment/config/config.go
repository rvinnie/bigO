package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	GRPC GRPCConfig
	GIN  GINConfig
}

type GRPCConfig struct {
	Port string `yaml:"port"`
}

type GINConfig struct {
	Mode string `yaml:"mode"`
}

func InitConfig(configDir string) (*Config, error) {
	viper.AddConfigPath(configDir)
	viper.SetConfigName("payment")
	var config Config

	if err := viper.UnmarshalKey("grpc", &config.GRPC); err != nil {
		return nil, err
	}

	initEnvVariables(&config)

	return &config, nil
}

func initEnvVariables(cfg *Config) {
	cfg.GIN.Mode = os.Getenv("GIN_MODE")
}
