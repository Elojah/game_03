package main

import (
	"context"
	"time"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/rtc"
	"github.com/elojah/game_03/pkg/rtc/dto"
	"github.com/pion/webrtc/v3"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) Connect(ctx context.Context, req *dto.SDP) (*dto.SDP, error) {
	logger := log.With().Str("method", "connect").Logger()

	if req == nil {
		return &dto.SDP{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// // #Authenticate
	u, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &dto.SDP{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// Prepare the configuration
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"}, // HD: set in config
			},
		},
	}

	// Create a new RTCPeerConnection
	pc, err := webrtc.NewPeerConnection(config)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create peer connection")

		return &dto.SDP{}, status.New(codes.Internal, err.Error()).Err()
	}

	// Set the handler for Peer connection state
	// This will notify you when the peer has connected/disconnected
	pc.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		logger.Info().Str("state", s.String()).Msg("peer connection state change")

		if s == webrtc.PeerConnectionStateFailed {
			// Wait until PeerConnection has had no network activity for 30 seconds or another failure. It may be reconnected using an ICE Restart.
			// Use webrtc.PeerConnectionStateDisconnected if you are interested in detecting faster timeout.
			// Note that the PeerConnection may come back from PeerConnectionStateDisconnected.
			// fmt.Println("Peer Connection has gone to failed exiting")
			// os.Exit(0)

			// close ?
			logger.Info().Str("state", s.String()).Msg("peer connection failed")

			return
		}
	})

	// Add handlers for setting up the connection.
	pc.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		logger.Info().Str("state", state.String()).Msg("ICE connection state change")
	})

	// Send the current time via a DataChannel to the remote peer every 3 seconds
	pc.OnDataChannel(func(d *webrtc.DataChannel) {
		d.OnOpen(func() {
			for range time.Tick(time.Second) {
				if err := d.SendText(time.Now().String()); err != nil {
					logger.Error().Err(err).Msg("failed to send text")

					continue
				}

				logger.Info().Msg("message sent")
			}
		})
	})
	// #Read and set remote description
	remoteSDP, err := req.MarshalRTC()
	if err != nil {
		logger.Error().Err(err).Msg("failed to read remote description")

		return &dto.SDP{}, status.New(codes.Internal, err.Error()).Err()
	}

	if err := pc.SetRemoteDescription(webrtc.SessionDescription(remoteSDP)); err != nil {
		logger.Error().Err(err).Msg("failed to set remote description")

		return &dto.SDP{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Set and write local description
	answer, err := pc.CreateAnswer(nil)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create answer")

		return &dto.SDP{}, status.New(codes.Internal, err.Error()).Err()
	}

	// Sets the LocalDescription, and starts our UDP listeners
	if err = pc.SetLocalDescription(answer); err != nil {
		logger.Error().Err(err).Msg("failed to set local description")

		return &dto.SDP{}, status.New(codes.Internal, err.Error()).Err()
	}

	if err := h.rtc.Insert(ctx, rtc.PC{
		ID:             u.ID,
		PeerConnection: pc,
	}); err != nil {
		logger.Error().Err(err).Msg("failed to create peer connection")

		return &dto.SDP{}, status.New(codes.Internal, err.Error()).Err()
	}

	var resp dto.SDP
	if err := resp.UnmarshalRTC(rtc.SDP(answer)); err != nil {
		logger.Error().Err(err).Msg("failed to read answer")

		return &dto.SDP{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &resp, nil
}
