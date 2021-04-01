package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/twitch"
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

	// Fetch twitch auth token
	token, err := h.twitch.GetToken(ctx, strings.TrimPrefix(req.Value, "oauth-session="), h.twitch.OAuth())
	if err != nil {
		logger.Error().Err(err).Msg("failed to get token")

		return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
	}

	// Fetch twitch user
	if err := h.twitch.GetUsers(
		ctx,
		twitch.Auth{
			Token:    token.AccessToken,
			ClientID: h.twitch.OAuth().ClientID,
		},
		twitch.UserFilter{},
		func(u twitch.User) error {
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
