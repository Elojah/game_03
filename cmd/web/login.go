package main

import (
	"net/http"

	"github.com/elojah/game_03/pkg/ulid"
	"github.com/rs/zerolog/log"
)

func (h handler) login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_ = ctx
	logger := log.With().Str("route", r.URL.EscapedPath()).Str("method", r.Method).Str("address", r.RemoteAddr).Logger()

	id := ulid.NewID()

	session, err := h.CookieStore.Get(r, "oauth-session")
	if err != nil {
		logger.Error().Err(err).Msg("failed to get oauth session")
	}

	session.AddFlash(id.String(), "oauth-state-callback")

	if err = session.Save(r, w); err != nil {
		logger.Error().Err(err).Msg("failed to save session")
		http.Error(w, "failed to save session", http.StatusInternalServerError)

		return
	}

	logger.Info().Msg("success")

	http.Redirect(w, r, h.Config.AuthCodeURL(id.String()), http.StatusTemporaryRedirect)
}
