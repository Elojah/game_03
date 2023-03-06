package main

import (
	"net/http"

	"github.com/gogo/protobuf/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h handler) refreshToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	logger := log.With().Str("route", r.URL.EscapedPath()).Str("method", r.Method).Str("address", r.RemoteAddr).Logger()

	// Refresh with auth.
	jwt, err := h.RefreshToken(ctx, &types.StringValue{Value: ""})
	if err != nil {
		logger.Error().Err(err).Msg("failed to refresh")

		if e, ok := status.FromError(err); ok {
			switch e.Code() { //nolint: exhaustive
			case codes.Internal:
				http.Error(w, e.Message(), http.StatusInternalServerError)
			case codes.InvalidArgument:
				http.Error(w, e.Message(), http.StatusBadRequest)
			default:
				http.Error(w, e.Message(), http.StatusInternalServerError)
			}
		} else {
			http.Error(w, "failed to signin", http.StatusInternalServerError)
		}

		return
	}

	// Set access token
	http.SetCookie(w, &http.Cookie{
		Name:     "access",
		Value:    jwt.AccessToken,
		Path:     "/",
		Secure:   true,
		HttpOnly: false,
		SameSite: http.SameSiteDefaultMode,
		Domain:   ".legacyfactory.com",
	})

	// Set refresh token
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh",
		Value:    jwt.RefreshToken,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
		Domain:   ".legacyfactory.com",
	})

	w.WriteHeader(http.StatusOK)

	logger.Info().Msg("success")
}