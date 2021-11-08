package main

import (
	"context"

	"github.com/elojah/go-grpc"
	"github.com/elojah/go-grpcweb"
	"github.com/elojah/go-http"
	"github.com/elojah/go-redis"
	"github.com/elojah/go-scylla"
	"github.com/ilyakaznacheev/cleanenv"
	"golang.org/x/oauth2"
)

type config struct {
	HTTP    http.Config    `json:"http"`
	GRPCWeb grpcweb.Config `json:"grpcweb"`
	GRPC    grpc.Config    `json:"grpc"`
	Scylla  scylla.Config  `json:"scylla"`
	Redis   redis.Config   `json:"redis"`
	Twitch  oauth2.Config  `json:"twitch"`
}

// Populate populates config object reading file and env.
func (c *config) Populate(ctx context.Context, filename string) error {
	return cleanenv.ReadConfig(filename, c)
}
