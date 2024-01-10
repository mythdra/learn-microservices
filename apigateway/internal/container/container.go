package container

import (
	"apigateway/internal/container/config"
	"context"
)

type IContainer interface {
	GetConfig() *config.Config
}

type container struct {
	config *config.Config
}

func NewContainer(ctx context.Context) (IContainer, error) {
	config, err := config.NewConfig(ctx)
	if err != nil {
		return nil, err
	}
	return &container{
		config: config,
	}, nil
}

func (c *container) GetConfig() *config.Config {
	return c.config
}
