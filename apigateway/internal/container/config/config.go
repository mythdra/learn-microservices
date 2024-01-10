package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	JWT *JWTConfig
}

func NewConfig(ctx context.Context) (*Config, error) {
	var c Config
	if err := envconfig.Process(ctx, &c); err != nil {
		return nil, err
	}
	return &c, nil
}
