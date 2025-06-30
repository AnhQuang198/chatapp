package config

import "github.com/spf13/viper"

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
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	var cfg Config
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
