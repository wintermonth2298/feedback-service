package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ServerPort string `json:"server_port"`
	SQL        SQL    `json:"sql"`
}

type SQL struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

func New(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %w", err)
	}

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("error decoding json: %w", err)
	}

	if err := file.Close(); err != nil {
		return nil, err
	}

	return &cfg, nil
}
