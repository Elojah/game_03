package main

import (
	"io"
	"net/http"

	"github.com/gogo/protobuf/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const maxTokenLength = 32768

func (h handler) signinGoogle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	logger := log.With().Str("route", r.URL.EscapedPath()).Str("method", r.Method).Str("address", r.RemoteAddr).Logger()

	// Fetch token from payload.
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error().Err(err).Msg("failed to read body")
		http.Error(w, "failed to read body", http.StatusInternalServerError)

		return
	}

	if len(raw) > maxTokenLength {
		logger.Error().Msg("invalid token format")
		http.Error(w, "invalid token format", http.StatusBadRequest)

		return
	}

	// Signin with auth.
	jwt, err := h.SigninGoogle(ctx, &types.StringValue{Value: string(raw)})
	if err != nil {
		logger.Error().Err(err).Msg("failed to signin")

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

	// Create JWT secure cookie.
	ck, err := h.cookie.Encode(ctx, "token", jwt.Value)
	if err != nil {
		logger.Error().Err(err).Msg("failed to encode token")
		http.Error(w, "failed to encode token", http.StatusInternalServerError)

		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    ck,
		Path:     "/",
		Secure:   true,
		HttpOnly: false,
		SameSite: http.SameSiteDefaultMode,
	})

	w.WriteHeader(http.StatusOK)

	logger.Info().Msg("success")
}
