package main

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func (h handler) redirect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := log.With().Str("route", r.URL.EscapedPath()).Str("method", r.Method).Str("address", r.RemoteAddr).Logger()

	session, err := h.CookieStore.Get(r, "oauth-session")
	if err != nil {
		logger.Error().Err(err).Msg("failed to get oauth session")
		http.Error(w, "failed to get oauth session", http.StatusInternalServerError)

		return
	}

	defer func() { _ = session.Save(r, w) }()

	challenge := session.Flashes("oauth-state-callback")
	state := r.FormValue("state")

	if len(challenge) < 1 && state == "" {
		http.Error(w, "missing state challenge", http.StatusBadRequest)

		return
	} else if challenge[0] != state {
		http.Error(w, "invalid oauth state", http.StatusBadRequest)

		return
	}

	token, err := h.Config.Exchange(ctx, r.FormValue("code"))
	if err != nil {
		http.Error(w, "invalid oauth state", http.StatusBadRequest)

		return
	}

	// add the oauth token to session
	session.Values["oauth-token"] = token

	logger.Info().Msg("success")

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
