package main

import (
	"context"

	"github.com/elojah/go-http"
	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	HTTP http.Config `json:"http"`
	Web  struct {
		Static string `json:"static"`
	} `json:"web"`
}

// Populate populates config object reading file and env.
func (c *config) Populate(ctx context.Context, filename string) error {
	return cleanenv.ReadConfig(filename, c)
}
