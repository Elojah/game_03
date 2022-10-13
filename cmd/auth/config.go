package main

import (
	"context"

	"github.com/elojah/go-grpc"
	"github.com/elojah/go-redis"
	"github.com/elojah/go-scylla"
	"github.com/ilyakaznacheev/cleanenv"
	"golang.org/x/oauth2"
)

type config struct {
	Scylla scylla.Config `json:"scylla"`
	GRPC   grpc.Config   `json:"grpc"`
	Redis  redis.Config  `json:"redis"`

	Twitch oauth2.Config `json:"twitch"`
	Google oauth2.Config `json:"google"`

	JWTSecret string `json:"jwt_secret"`
}

// Populate populates config object reading file and env.
func (c *config) Populate(ctx context.Context, filename string) error {
	return cleanenv.ReadConfig(filename, c)
}
