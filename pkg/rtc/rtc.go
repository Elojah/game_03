package rtc

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
	"github.com/pion/webrtc/v3"
)

type SDP webrtc.SessionDescription

type ICECandidate webrtc.ICECandidate

type ICECandidateInit webrtc.ICECandidateInit

type PC struct {
	*webrtc.PeerConnection

	ID ulid.ID
}

type Filter struct {
	ID ulid.ID
}

type Store interface {
	Insert(context.Context, PC) error
	Fetch(context.Context, Filter) (PC, error)
	Delete(context.Context, Filter) error
}

type App interface {
	Store
}
