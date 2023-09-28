package main

import (
	"github.com/elojah/game_03/cmd/core/grpc"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/pbtypes"
	"github.com/elojah/game_03/pkg/rtc"
	"github.com/elojah/game_03/pkg/rtc/dto"
	"github.com/pion/webrtc/v3"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) ReceiveICE(req *pbtypes.Empty, stream grpc.Core_ReceiveICEServer) error {
	ctx := stream.Context()
	logger := log.With().Str("method", "receive_ice").Logger()

	if req == nil {
		return status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// // #Authenticate
	u, err := h.user.Auth(ctx, "access")
	if err != nil {
		return status.New(codes.Unauthenticated, err.Error()).Err()
	}

	pc, err := h.rtc.Fetch(ctx, rtc.Filter{ID: u.ID})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch peer connection")

		return status.New(codes.NotFound, err.Error()).Err()
	}

	ch := make(chan dto.ICECandidate)

	pc.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate == nil {
			return
		}

		logger.Info().Str("candidate", candidate.String()).Msg("ICE candidate received")

		var result dto.ICECandidate
		if err := result.UnmarshalRTC(rtc.ICECandidate(*candidate)); err != nil {
			logger.Error().Err(err).Str("candidate", candidate.String()).Msg("failed to read ice candidate")

			return
		}

		ch <- result
	})

	logger.Info().Msg("sending ice candidates...")

	for {
		select {
		case _ = <-ctx.Done():
			logger.Info().Msg("done")

			return ctx.Err()
		case ic := <-ch:
			if err := stream.SendMsg(&ic); err != nil {
				logger.Error().Err(err).Msg("failed to send message")

				return status.New(codes.Internal, err.Error()).Err()
			}

			logger.Info().Msg("sent ice candidate")
		}
	}
}
