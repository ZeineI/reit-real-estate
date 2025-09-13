package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

const configFile = "./config/config.yaml"

type Config struct {
	AppHost        string         `json:"app_host"`
	AppPort        string         `json:"app_port"`
	DatabaseConfig DatabaseConfig `yaml:"db"`
}

type DatabaseConfig struct {
	Host               string `yaml:"host"`
	Port               int    `yaml:"port"`
	User               string `yaml:"user"`
	Password           string `yaml:"password"`
	DatabaseName       string `yaml:"database_name"`
	MaxOpenConnections int    `yaml:"max_open_connections"`
	MaxIdleConnections int    `yaml:"max_idle_connections"`
}

func LoadConfig() (*Config, error) {
	config := &Config{}
	rawYaml, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("config.ReadFile error: %w", err)
	}

	err = yaml.Unmarshal(rawYaml, &config)
	if err != nil {
		return nil, fmt.Errorf("config.Unmarshal error: %w", err)
	}
	return config, nil
}
