package config

import "github.com/spf13/viper"

type Config struct {
	GRPC GRPCConfig
}

type GRPCConfig struct {
	Port string `yaml:"port"`
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

	return &cfg, nil
}
