package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	LogLevel     zapcore.Level `env:"LOG_LEVEL" default="Info"`
	ServiceName  string        `env:"SERVICE_NAME"`
	Environment  string        `env:"ENV"`
	ProjectID    string        `env:"GCP_PROJECT_ID"`
	InstanceName string        `env:"INSTANCE_NAME"`
	DatabaseName string        `env:"DATABASE_NAME"`
}

//New returns type of Config
func New(ctx context.Context) (*Config, error) {
	var c Config
	err := envconfig.Process(ctx, &c)
	if err != nil {
		return &Config{}, err
	}

	return &c, nil
}
