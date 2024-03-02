package config

import (
	"AlekseyPromet/simple_sso/internal/models"
	"fmt"

	uber "go.uber.org/config"
)

func NewConfig() (models.Config, error) {
	cfg := models.Config{}

	loader, err := uber.NewYAML(uber.File("config.yaml"))
	if err != nil {
		return cfg, fmt.Errorf("confg loader failed: %w", err)
	}

	if err := loader.Get("app").Populate(&cfg); err != nil {
		return cfg, fmt.Errorf("confg loader failed: %w", err)
	}

	return cfg, nil
}
