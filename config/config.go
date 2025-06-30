package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Config struct {
	//load config server
	Server struct {
		Port string
	}

	//load config for database
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		Driver   string
		SSLMode  string
	}
}

func LoadConfig(path string) (*Config, error) {
	filePath := fmt.Sprintf("config/%s.yaml", path)
	content, err := os.ReadFile(filePath) //load all properties from yaml
	if err != nil {
		return nil, fmt.Errorf("read config file error: %w", err)
	}

	expanded := os.ExpandEnv(string(content)) //replace value from environment

	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(strings.NewReader(expanded)); err != nil {
		return nil, fmt.Errorf("viper read config error: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("viper unmarshal error: %w", err)
	}

	return &cfg, nil
}
