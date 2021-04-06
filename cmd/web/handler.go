package main

import (
	"context"

	"github.com/elojah/game_03/cmd/auth/grpc"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/twitch"
)

type handler struct {
	oauth2.Config
	*sessions.CookieStore

	grpc.AuthClient
}

func (h *handler) Dial(ctx context.Context, w web) error {
	h.Config = w.Twitch
	h.Config.Endpoint = twitch.Endpoint

	h.CookieStore = sessions.NewCookieStore([]byte(w.Secret))

	return nil
}

func (h *handler) Close(ctx context.Context) error {
	return nil
}
