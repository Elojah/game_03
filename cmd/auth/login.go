package main

import (
	"context"
	"errors"

	gerrors "github.com/elojah/game_03/pkg/errors"
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
		return &types.StringValue{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
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

		// return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Check if user exist
	u, err := h.user.Fetch(ctx, user.Filter{TwitchID: &tu.ID})
	if err != nil {
		if errors.As(err, &gerrors.ErrNotFound{}) {
			// #Create user
			u = user.U{
				ID:       ulid.NewID(),
				TwitchID: tu.ID,
			}
			if err := h.user.Insert(ctx, u); err != nil {
				logger.Error().Err(err).Msg("failed to create user")

				return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
			}
		} else {
			logger.Error().Err(err).Msg("failed to fetch user")

			return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
		}
	}

	// Create session
	ses := user.Session{
		ID:          ulid.NewID(),
		UserID:      u.ID,
		TwitchToken: req.Value,
	}
	if err := h.user.InsertSession(ctx, ses); err != nil {
		logger.Error().Err(err).Msg("failed to create session")

		return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &types.StringValue{Value: ses.ID.String()}, nil
}
