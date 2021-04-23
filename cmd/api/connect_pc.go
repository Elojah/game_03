package main

import (
	"github.com/elojah/game_03/cmd/api/grpc"
	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) ConnectPC(req *entity.PC, stream grpc.API_ConnectPCServer) error {
	ctx := stream.Context()
	logger := log.With().Str("method", "create_pc").Logger()

	if req == nil {
		return status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	_, err := h.user.Auth(ctx)
	if err != nil {
		return status.New(codes.Unauthenticated, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return nil
}
