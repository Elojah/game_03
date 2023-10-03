package main

import (
	"context"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/pbtypes"
	"github.com/elojah/game_03/pkg/rtc"
	"github.com/elojah/game_03/pkg/rtc/dto"
	"github.com/pion/webrtc/v3"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) SendICE(ctx context.Context, req *dto.ICECandidate) (*pbtypes.Empty, error) {
	logger := log.With().Str("method", "send_ice").Logger()

	if req == nil {
		return &pbtypes.Empty{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	u, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &pbtypes.Empty{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	pc, err := h.rtc.Fetch(ctx, rtc.Filter{ID: u.ID})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch peer connection")

		return &pbtypes.Empty{}, status.New(codes.NotFound, err.Error()).Err()
	}

	ic, err := req.MarshalRTC()
	if err != nil {
		logger.Error().Err(err).Msg("failed to read candidate")

		return &pbtypes.Empty{}, status.New(codes.Internal, err.Error()).Err()
	}

	if err := pc.AddICECandidate(webrtc.ICECandidateInit(ic)); err != nil {
		logger.Error().Err(err).Msg("failed to add candidate")

		return &pbtypes.Empty{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Str("candidate", ic.Candidate).Msg("candidate added")

	return &pbtypes.Empty{}, nil
}
