package main

import (
	"context"

	"github.com/elojah/game_03/cmd/api/grpc"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type updatePCSender struct {
	grpc.API_UpdatePCClient
}

func (h *handler) UpdatePC(ctx context.Context) (grpc.API_UpdatePCClient, error) {
	logger := log.With().Str("method", "update_pc").Logger()

	// #Authenticate
	ses, err := h.user.Auth(ctx)
	if err != nil {
		return nil, status.New(codes.Unauthenticated, err.Error()).Err()
	}
}

// func (s updatePCSender)
