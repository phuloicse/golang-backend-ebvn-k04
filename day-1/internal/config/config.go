package config

import (
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port        string `envconfig:"APP_PORT" default:"8081"`
	InstanceID  string `envconfig:"INSTANCE_ID"`
	ServiceName string `envconfig:"SERVICE_NAME" default:"golang-backend-ebvn-k04-day-1"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("api", cfg)
	if err != nil {
		return nil, err
	}
	if cfg.InstanceID == "" {
		cfg.InstanceID = uuid.NewString()
	}

	return cfg, err
}
