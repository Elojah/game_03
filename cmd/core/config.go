package main

import (
	"context"

	"github.com/elojah/game_03/pkg/user"
	"github.com/elojah/go-http"
	"github.com/elojah/go-redis"
	"github.com/elojah/go-scylla"
	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	HTTP   http.Config   `json:"http"`
	Scylla scylla.Config `json:"scylla"`
	Redis  redis.Config  `json:"redis"`

	Session user.ConfigSession `json:"session"`
}

// Populate populates config object reading file and env.
func (c *config) Populate(ctx context.Context, filename string) error {
	return cleanenv.ReadConfig(filename, c)
}
