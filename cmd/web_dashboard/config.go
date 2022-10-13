package main

import (
	"context"

	"github.com/elojah/go-grpc"
	"github.com/elojah/go-http"
	"github.com/elojah/go-redis"
	"github.com/ilyakaznacheev/cleanenv"
)

type web struct {
	Static string `json:"static"`
}

type config struct {
	HTTP       http.Config       `json:"http"`
	Web        web               `json:"web"`
	Redis      redis.Config      `json:"redis"`
	AuthClient grpc.ConfigClient `json:"auth_client"`
}

// Populate populates config object reading file and env.
func (c *config) Populate(ctx context.Context, filename string) error {
	return cleanenv.ReadConfig(filename, c)
}
