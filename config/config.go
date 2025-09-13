package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

const configFile = "/config/config.yaml"

type Config struct {
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
