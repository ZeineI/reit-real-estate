package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

const configFile = "./config/config.yaml"

type Config struct {
	AppConfig      AppConfig      `yaml:"app"`
	DatabaseConfig DatabaseConfig `yaml:"db"`
	SolanaConfig   SolanaConfig   `yaml:"solana"`
}

type AppConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
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

type SolanaConfig struct {
	RpcURL       string `yaml:"rpc_url"`
	ProgramID    string `yaml:"program_id"`
	TokenAddress string `yaml:"token_address"`
	ReitMint     string `yaml:"reit_mint"`
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
