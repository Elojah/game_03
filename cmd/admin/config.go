package main

import (
	"context"

	"github.com/elojah/go-grpc"
	"github.com/elojah/go-redis"
	"github.com/elojah/go-scylla"
	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Scylla scylla.Config `json:"scylla"`
	Redis  redis.Config  `json:"redis"`
	GRPC   grpc.Config   `json:"grpc"`
}

// Populate populates config object reading file and env.
func (c *config) Populate(ctx context.Context, filename string) error {
	return cleanenv.ReadConfig(filename, c)
}
