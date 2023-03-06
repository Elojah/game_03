package main

import (
	"context"

	"github.com/elojah/game_03/pkg/user"
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

	Session user.ConfigSession `json:"session"`

	Twitch oauth2.Config `json:"twitch"`
	Google oauth2.Config `json:"google"`
}

// Populate populates config object reading file and env.
func (c *config) Populate(ctx context.Context, filename string) error {
	return cleanenv.ReadConfig(filename, c)
}
