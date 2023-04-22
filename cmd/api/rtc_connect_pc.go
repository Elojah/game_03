package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/user/dto"
	"github.com/gogo/protobuf/types"
	"github.com/pion/webrtc/v3"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) RTCConnectPC(ctx context.Context, req *dto.SDP) (*types.Empty, error) {
	logger := log.With().Str("method", "rtc_connect_pc").Logger()

	if req == nil {
		return &types.Empty{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// // #Authenticate
	// u, err := h.user.Auth(ctx, "access")
	// if err != nil {
	// 	return &types.Empty{}, status.New(codes.Unauthenticated, err.Error()).Err()
	// }

	// _ = u

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
		log.Error().Err(err).Msg("failed to create peer connection")

		return &types.Empty{}, status.New(codes.Internal, err.Error()).Err()
	}

	defer func() {
		if cErr := pc.Close(); cErr != nil {
			log.Error().Err(err).Msg("failed to close peer connection")
		}
	}()

	// Set the handler for Peer connection state
	// This will notify you when the peer has connected/disconnected
	pc.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		log.Info().Str("state", s.String()).Msg("peer connection state change")

		if s == webrtc.PeerConnectionStateFailed {
			// Wait until PeerConnection has had no network activity for 30 seconds or another failure. It may be reconnected using an ICE Restart.
			// Use webrtc.PeerConnectionStateDisconnected if you are interested in detecting faster timeout.
			// Note that the PeerConnection may come back from PeerConnectionStateDisconnected.
			// fmt.Println("Peer Connection has gone to failed exiting")
			// os.Exit(0)

			// HD: not sure we want to exit server here...
			log.Info().Str("state", s.String()).Msg("peer connection failed")

			return
		}
	})

	// Register data channel creation handling
	pc.OnDataChannel(func(d *webrtc.DataChannel) {
		log.Info().Str("label", d.Label()).Uint16("id", *d.ID()).Msg("data channel created")

		// Register channel opening handling
		d.OnOpen(func() {
			for range time.NewTicker(5 * time.Second).C {
				message := "test"

				// Send the message as text
				log.Info().Msg("send text message")

				if err := d.SendText(message); err != nil {
					log.Error().Err(err).Msg("failed to send text message")

					return
				}
			}
		})

		// Register text message handling
		d.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Printf("Message from DataChannel '%s': '%s'\n", d.Label(), string(msg.Data))
		})
	})

	// Add handlers for setting up the connection.
	pc.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {})
	pc.OnICECandidate(func(candidate *webrtc.ICECandidate) {})

	var remoteSDP webrtc.SessionDescription
	b, err := base64.StdEncoding.DecodeString(req.Encoded)
	if err != nil {
		log.Error().Err(err).Msg("failed to decode remote sdp")

		return &types.Empty{}, status.New(codes.Internal, err.Error()).Err()
	}

	if err := json.Unmarshal(b, &remoteSDP); err != nil {
		log.Error().Err(err).Msg("failed to set read remote sdp")

		return &types.Empty{}, status.New(codes.Internal, err.Error()).Err()
	}

	if err := pc.SetRemoteDescription(remoteSDP); err != nil {
		log.Error().Err(err).Msg("failed to set remote description")

		return &types.Empty{}, status.New(codes.Internal, err.Error()).Err()
	}

	// Create an answer
	answer, err := pc.CreateAnswer(nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to create answer")

		return &types.Empty{}, status.New(codes.Internal, err.Error()).Err()
	}

	// Create channel that is blocked until ICE Gathering is complete
	complete := webrtc.GatheringCompletePromise(pc)

	// Sets the LocalDescription, and starts our UDP listeners
	if err = pc.SetLocalDescription(answer); err != nil {
		log.Error().Err(err).Msg("failed to set local description")

		return &types.Empty{}, status.New(codes.Internal, err.Error()).Err()
	}

	// Block until ICE Gathering is complete, disabling trickle ICE
	// we do this because we only can exchange one signaling message
	// in a production application you should exchange ICE Candidates via OnICECandidate
	<-complete

	logger.Info().Msg("success")

	return &types.Empty{}, nil
}
