package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	GRPC   GRPCConfig
	OpenAI OpenAIConfig
}

type GRPCConfig struct {
	Port string `yaml:"port"`
}

type OpenAIConfig struct {
	Key string
}

func InitConfig(configDir string) (*Config, error) {
	viper.AddConfigPath(configDir)
	viper.SetConfigName("algorithm-complexity")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.UnmarshalKey("gRPC", &cfg.GRPC); err != nil {
		return nil, err
	}

	setEnvVariables(&cfg)

	return &cfg, nil
}

func setEnvVariables(cfg *Config) {
	cfg.OpenAI.Key = os.Getenv("OPEN_API_KEY")
}
