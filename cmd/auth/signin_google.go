package main

import (
	"context"
	"errors"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/elojah/game_03/pkg/user"
	"github.com/gogo/protobuf/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) SigninGoogle(ctx context.Context, req *types.StringValue) (*types.StringValue, error) {
	logger := log.With().Str("method", "signin_google").Logger()

	if req == nil {
		logger.Error().Err(gerrors.ErrNullRequest{}).Msg("null request")

		return &types.StringValue{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Validate token
	gid, err := h.google.Signin(ctx, req.Value)
	if err != nil {
		logger.Error().Err(err).Msg("failed to validate token")

		return &types.StringValue{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	// #Check if user exist
	u, err := h.user.Fetch(ctx, user.Filter{GoogleID: &gid})
	if err != nil {
		if errors.As(err, &gerrors.ErrNotFound{}) {
			// #Create user
			u = user.U{
				ID:       ulid.NewID(),
				GoogleID: gid,
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

	// #Create JWT
	jwt, err := h.user.CreateJWT(ctx, u)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create JWT")

		return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &types.StringValue{Value: jwt}, nil
}
