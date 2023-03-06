package main

import (
	"context"
	"errors"

	gerrors "github.com/elojah/game_03/pkg/errors"
	gtwitch "github.com/elojah/game_03/pkg/twitch"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/elojah/game_03/pkg/user"
	"github.com/elojah/game_03/pkg/user/dto"
	"github.com/gogo/protobuf/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) SigninTwitch(ctx context.Context, req *types.StringValue) (*dto.SigninResp, error) {
	logger := log.With().Str("method", "signin_twitch").Logger()

	if req == nil {
		return &dto.SigninResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
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

		return &dto.SigninResp{}, status.New(codes.Internal, err.Error()).Err()
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

				return &dto.SigninResp{}, status.New(codes.Internal, err.Error()).Err()
			}
		} else {
			logger.Error().Err(err).Msg("failed to fetch user")

			return &dto.SigninResp{}, status.New(codes.Internal, err.Error()).Err()
		}
	}

	// #Create JWT
	jwt, err := h.user.CreateJWT(ctx, u)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create JWT")

		return &dto.SigninResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &dto.SigninResp{
		AccessToken: jwt,
	}, nil
}
