package main

import (
	"context"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/twitch"
)

type handler struct {
	oauth2.Config
	*sessions.CookieStore
}

func (h *handler) Dial(ctx context.Context, w web) error {
	h.Config = oauth2.Config{
		ClientID:     w.Twitch.IDClient,
		ClientSecret: w.Twitch.SecretToken,
		Scopes:       w.Twitch.Scopes,
		Endpoint:     twitch.Endpoint,
		RedirectURL:  w.Twitch.RedirectURL,
	}

	h.CookieStore = sessions.NewCookieStore([]byte(w.Secret))

	return nil
}

func (h *handler) Close(ctx context.Context) error {
	return nil
}
