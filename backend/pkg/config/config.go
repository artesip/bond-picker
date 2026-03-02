package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(path string) *Config {
	file, err := os.ReadFile(path)

	if err != nil {
		panic(fmt.Errorf("config load error: %w", err.Error()))
	}

	cfg := new(Config)

	err = yaml.Unmarshal(file, cfg)
	if err != nil {
		panic(fmt.Errorf("config unmarshal error: %w", err.Error()))
	}

	return cfg
}
