package main

import (
	"context"

	"github.com/elojah/game_03/pkg/errors"
	gtwitch "github.com/elojah/game_03/pkg/twitch"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/elojah/game_03/pkg/user"
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
	var tu gtwitch.User

	if err := h.twitch.GetUsers(
		ctx,
		gtwitch.Auth{
			Token:    req.Value,
			ClientID: h.twitch.OAuth().ClientID,
		},
		gtwitch.UserFilter{},
		func(u gtwitch.User) error {
			tu = u

			return nil
		},
	); err != nil {
		logger.Error().Err(err).Msg("failed to get users")

		return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
	}

	// Create new session
	id := ulid.NewID()

	_, err := h.user.Fetch(ctx, user.Filter{TwitchID: &tu.ID})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch user")

		return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &types.StringValue{Value: id.String()}, nil
}
