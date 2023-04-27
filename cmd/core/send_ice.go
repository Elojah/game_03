package main

import (
	"io"

	ggrpc "github.com/elojah/game_03/cmd/core/grpc"
	"github.com/elojah/game_03/pkg/rtc"
	"github.com/gogo/protobuf/types"
	"github.com/pion/webrtc/v3"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) SendICE(stream ggrpc.Core_SendICEServer) error {
	ctx := stream.Context()
	logger := log.With().Str("method", "send_ice").Logger()

	// #Authenticate

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

	for {
		// Check context done
		select {
		case _ = <-ctx.Done():
			logger.Info().Msg("done")

			return ctx.Err()
		default:
		}

		// Receive candidate
		candidate, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&types.Empty{})
		}

		if err != nil {
			logger.Error().Err(err).Msg("failed to receive candidate")

			return status.New(codes.Internal, err.Error()).Err()
		}

		if candidate == nil {
			continue
		}

		ic, err := candidate.MarshalRTC()
		if err != nil {
			logger.Error().Err(err).Msg("failed to read candidate")

			continue
		}

		if err := pc.AddICECandidate(webrtc.ICECandidateInit(ic)); err != nil {
			logger.Error().Err(err).Msg("failed to add candidate")

			continue
		}

		logger.Info().Str("candidate", ic.Candidate).Msg("candidate added")
	}
}
