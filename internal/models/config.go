package models

import "time"

type Config struct {
	Env       TypeEnv       `yaml:"env,omitempty" env-default:"local"`
	StorgeDsn string        `yaml:"storge_dsn,required"`
	TokenTTL  time.Duration `yaml:"token_ttl,required"`
	Grpc
}

type Grpc struct {
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}
