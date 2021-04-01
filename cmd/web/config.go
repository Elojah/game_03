package main

import (
	"context"

	"github.com/elojah/go-http"
	"github.com/ilyakaznacheev/cleanenv"
	"golang.org/x/oauth2"
)

type web struct {
	Static string        `json:"static"`
	Secret string        `json:"secret"`
	Twitch oauth2.Config `json:"twitch"`
}

type config struct {
	HTTP http.Config `json:"http"`
	Web  web         `json:"web"`
}

// Populate populates config object reading file and env.
func (c *config) Populate(ctx context.Context, filename string) error {
	return cleanenv.ReadConfig(filename, c)
}
