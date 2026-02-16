package config

import (
	"backend/internal/domain"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(path string) *domain.Config {
	file, err := os.ReadFile(path)

	if err != nil {
		panic(fmt.Errorf("config load error: %w", err.Error()))
	}

	cfg := new(domain.Config)

	err = yaml.Unmarshal(file, cfg)
	if err != nil {
		panic(fmt.Errorf("config unmarshal error: %w", err.Error()))
	}

	return cfg
}
