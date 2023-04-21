package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/elojah/game_03/pkg/ulid"
	glog "github.com/elojah/go-log"
	"github.com/hashicorp/go-multierror"
	"github.com/pion/sdp/v3"
	"github.com/pion/webrtc/v3"
	"github.com/rs/zerolog/log"
)

const (
	// Time allocated for init phase (connections + setup).
	initTO = 30 * time.Second
)

var version string

type closer interface {
	Close(context.Context) error
}

type closers []closer

func (cs closers) Close(ctx context.Context) error {
	var result *multierror.Error

	for _, c := range cs {
		if c != nil {
			result = multierror.Append(result, c.Close(ctx))
		}
	}

	return result.ErrorOrNil()
}

// run services.
func run(prog string, filename string) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, initTO)
	// no need for defer cancel
	_ = cancel

	var cs []closer

	logs := glog.Service{}
	if err := logs.Dial(ctx, glog.Config{}); err != nil {
		fmt.Println("failed to dial logger")

		return
	}

	cs = append(cs, &logs)

	log.Logger = logs.With().Str("version", version).Str("exe", prog).Logger()

	// read config
	cfg := config{}
	if err := cfg.Populate(ctx, filename); err != nil {
		log.Error().Err(err).Msg("failed to read config")

		return
	}

	/*
	 */

	// Prepare the configuration
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"}, // HD: set in config
			},
		},
	}

	// Create a new RTCPeerConnection
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Error().Err(err).Msg("failed to create peer connection")

		return
	}

	defer func() {
		if cErr := peerConnection.Close(); cErr != nil {
			log.Error().Err(err).Msg("failed to close peer connection")
		}
	}()

	// Set the handler for Peer connection state
	// This will notify you when the peer has connected/disconnected
	peerConnection.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		log.Info().Str("state", s.String()).Msg("peer connection state change")

		if s == webrtc.PeerConnectionStateFailed {
			// Wait until PeerConnection has had no network activity for 30 seconds or another failure. It may be reconnected using an ICE Restart.
			// Use webrtc.PeerConnectionStateDisconnected if you are interested in detecting faster timeout.
			// Note that the PeerConnection may come back from PeerConnectionStateDisconnected.

			// HD: not sure we want to exit server here...

			// fmt.Println("Peer Connection has gone to failed exiting")
			// os.Exit(0)
		}
	})

	// Register data channel creation handling
	peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
		log.Info().Str("label", d.Label()).Uint16("id", *d.ID()).Msg("created data channel")

		// Register channel opening handling
		d.OnOpen(func() {
			for range time.NewTicker(5 * time.Second).C {
				message := "test"

				// Send the message as text
				sendErr := d.SendText(message)
				if sendErr != nil {
					panic(sendErr)
				}
			}
		})

		// Register text message handling
		d.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Printf("Message from DataChannel '%s': '%s'\n", d.Label(), string(msg.Data))
		})
	})

	// Set the remote SessionDescription

	local, err := url.Parse("localhost:4283")
	if err != nil {
		log.Error().Err(err).Msg("failed to parse local URL")

		return
	}

	gsdp := sdp.SessionDescription{
		Version: 0,
		Origin: sdp.Origin{
			Username:       ulid.NewID().String(),
			SessionID:      0,
			SessionVersion: 0,
			NetworkType:    "IN",
			AddressType:    "IP4",
			UnicastAddress: "127.0.0.1",
		},
		SessionName:        sdp.SessionName(ulid.NewID().String()),
		SessionInformation: func(s sdp.Information) *sdp.Information { return &s }(""),
		URI:                local,
		EmailAddress:       func(s sdp.EmailAddress) *sdp.EmailAddress { return &s }("admin@legacyfactory.com"),
		PhoneNumber:        func(s sdp.PhoneNumber) *sdp.PhoneNumber { return &s }("+33000000000"),
		ConnectionInformation: &sdp.ConnectionInformation{
			NetworkType: "IN",
			AddressType: "IP4", // "ip6" ?
			Address: func(address string, ttl int, r int) *sdp.Address {
				return &sdp.Address{
					Address: address,
					TTL:     nil,
					Range:   &r,
				}
			}("127.0.0.1", 0, 1), // probably wrong ?
		},
		Bandwidth:        nil,
		TimeDescriptions: []sdp.TimeDescription{{Timing: sdp.Timing{StartTime: 0, StopTime: 0}}},
		TimeZones:        nil, // to change probably
		EncryptionKey:    nil, // deprecated
		Attributes: []sdp.Attribute{
			{Key: "entity", Value: "{protoencoded ?}"},
		},
		MediaDescriptions: []*sdp.MediaDescription{{
			MediaName: sdp.MediaName{
				Media: "application",
				Port: sdp.RangedPort{
					Value: 4283,
					Range: func(n int) *int { return &n }(1),
				},
				Protos:  []string{"UDP"},
				Formats: []string{"custom"},
			},
			Attributes: []sdp.Attribute{
				{Key: "mid", Value: "application"},
				{Key: "ice-ufrag", Value: ""},
				{Key: "ice-pwd", Value: ""},
				{Key: "fingerprint", Value: ulid.NewID().String() + " basic"},
			},
		}},
	}

	raw, err := gsdp.Marshal()
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal description")

		return
	}

	if err := peerConnection.SetRemoteDescription(webrtc.SessionDescription{
		Type: webrtc.SDPTypeOffer,
		SDP:  string(raw),
	}); err != nil {
		log.Error().Err(err).Msg("failed to set remote description")

		return
	}

	// Create an answer
	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to create answer")

		return
	}

	// Create channel that is blocked until ICE Gathering is complete
	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	// Sets the LocalDescription, and starts our UDP listeners
	err = peerConnection.SetLocalDescription(answer)
	if err != nil {
		log.Error().Err(err).Msg("failed to set local description")

		return
	}

	// Block until ICE Gathering is complete, disabling trickle ICE
	// we do this because we only can exchange one signaling message
	// in a production application you should exchange ICE Candidates via OnICECandidate
	<-gatherComplete

	log.Info().Msg("core up")

	// listen for signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	for sig := range c {
		switch sig {
		case syscall.SIGHUP:
			fallthrough
		case syscall.SIGINT:
			fallthrough
		case syscall.SIGTERM:
			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, initTO)

			defer cancel()

			if err := closers(cs).Close(ctx); err != nil {
				fmt.Printf("error closing service: %s\n", err.Error())

				return
			}

			fmt.Println("successfully closed service")

			return
		}
	}
}

func main() {
	args := os.Args
	if len(args) != 2 { //nolint: gomnd
		fmt.Printf("Usage: ./%s configfile\n", args[0])

		return
	}

	run(args[0], args[1])
}
