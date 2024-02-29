package models

import (
	"time"

	"go.uber.org/config"
	"go.uber.org/fx"
)

type Config struct {
	Env               string        `yaml:"env" default:"local"`
	Port              string        `yaml:"port" default:"8080"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout" default:"10s"`
}

type JWTConfig struct {
	TokenTTL time.Duration `yaml:"token_ttl" default:"30m"`
}

type GrpcConfig struct {
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type StorageConfig struct {
	StorgeDsn  string `yaml:"dsn"`
	Collection string `yaml:"collection"`
	Database   string `yaml:"database"`
}

type ResultConfig struct {
	fx.Out
	Provider config.Provider
	Config   *Config
}
