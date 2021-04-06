package main

import (
	"context"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/twitch"
	"github.com/elojah/game_03/pkg/twitch/dto"
	"github.com/elojah/game_03/pkg/user"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) ListFollow(ctx context.Context, req *dto.ListFollowReq) (*dto.ListFollowResp, error) {
	logger := log.With().Str("method", "list_follow").Logger()

	if req == nil {
		return &dto.ListFollowResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	ses, err := h.user.Auth(ctx)
	if err != nil {
		return &dto.ListFollowResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// #Fetch user
	u, err := h.user.Fetch(ctx, user.Filter{ID: &ses.UserID})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch user")

		return &dto.ListFollowResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Fetch follows
	var result dto.ListFollowResp

	cursor, err := h.twitch.GetFollows(ctx,
		twitch.Auth{
			Token:    ses.TwitchToken,
			ClientID: h.twitch.OAuth().ClientID,
		},
		twitch.FollowFilter{
			FromID: &u.TwitchID,
			After:  &req.After,
			First:  &req.First,
		},
		func(fo twitch.Follow) error {
			result.Follows = append(result.Follows, fo)

			return nil
		},
	)
	result.Cursor = cursor.Cursor
	result.Total = uint64(cursor.Total)

	logger.Info().Msg("success")

	return &result, nil
}
