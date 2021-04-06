package main

import (
	"net/http"

	"github.com/gogo/protobuf/types"
	"github.com/rs/zerolog/log"
)

func (h handler) redirect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := log.With().Str("route", r.URL.EscapedPath()).Str("method", r.Method).Str("address", r.RemoteAddr).Logger()

	// #Fetch session
	session, err := h.CookieStore.Get(r, "oauth-session")
	if err != nil {
		logger.Error().Err(err).Msg("failed to get oauth session")
		http.Error(w, "failed to get oauth session", http.StatusInternalServerError)

		return
	}

	// #Check session
	challenge := session.Flashes("oauth-state-callback")
	state := r.FormValue("state")

	if len(challenge) < 1 || state == "" {
		http.Error(w, "missing state challenge", http.StatusBadRequest)

		return
	} else if challenge[0] != state {
		http.Error(w, "invalid oauth state", http.StatusBadRequest)

		return
	}

	// #Fetch token
	token, err := h.Config.Exchange(ctx, r.FormValue("code"))
	if err != nil {
		logger.Error().Err(err).Msg("invalid oauth state")
		http.Error(w, "invalid oauth state", http.StatusBadRequest)

		return
	}

	// #Fetch game token
	at, err := h.AuthClient.Login(ctx, &types.StringValue{Value: token.AccessToken})
	if err != nil {
		logger.Error().Err(err).Msg("failed to login")
		http.Error(w, "failed to login", http.StatusInternalServerError)

		return
	}

	// #Delete session
	session.Options.MaxAge = -1

	if err := session.Save(r, w); err != nil {
		logger.Error().Err(err).Msg("failed to save session")
		http.Error(w, "failed to save session", http.StatusInternalServerError)

		return
	}

	// #Save game token client-side
	if at != nil {
		http.SetCookie(w, &http.Cookie{
			Name:     "auth-token",
			Value:    at.Value,
			SameSite: http.SameSiteStrictMode,
			Secure:   true,
			MaxAge:   3600, // nolint: gomnd
		})
	}

	logger.Info().Msg("success")

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
