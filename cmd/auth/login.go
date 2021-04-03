package main

import (
	"context"
	"fmt"

	"github.com/elojah/game_03/pkg/errors"
	gtwitch "github.com/elojah/game_03/pkg/twitch"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/gogo/protobuf/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) Login(ctx context.Context, req *types.StringValue) (*types.StringValue, error) {
	logger := log.With().Str("method", "login").Logger()

	if req == nil {
		return &types.StringValue{}, status.New(codes.Internal, errors.ErrNullRequest{}.Error()).Err()
	}

	// Fetch twitch user
	if err := h.twitch.GetUsers(
		ctx,
		gtwitch.Auth{
			Token:    req.Value,
			ClientID: h.twitch.OAuth().ClientID,
		},
		gtwitch.UserFilter{},
		func(u gtwitch.User) error {
			fmt.Println(u)

			return nil
		},
	); err != nil {
		logger.Error().Err(err).Msg("failed to get users")

		return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
	}

	// Create new session
	id := ulid.NewID()

	// if err := h.user.UpsertSession(ctx, user.Session{ID: id}); err != nil {
	// 	logger.Error().Err(err).Msg("failed to create new session")

	// 	return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
	// }

	logger.Info().Msg("success")

	return &types.StringValue{Value: id.String()}, nil
}
