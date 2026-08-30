package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("config load error: %w", err)
	}

	cfg := new(Config)

	err = yaml.Unmarshal(file, cfg)
	if err != nil {
		return nil, fmt.Errorf("yaml config unmarshal error: %w", err)
	}

	return cfg, nil
}
