package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App  `yaml:"app"`
	HTTP `yaml:"http"`
}

type App struct {
	ProjectName string `yaml:"project_name"`
	Version     string `yaml:"version"`
}

type HTTP struct {
	Port int `yaml:"port"`
}

func LoadConfig(path string) (Config, error) {
	var cfg Config

	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return cfg, fmt.Errorf("failed to read config: %w", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return cfg, nil
}
