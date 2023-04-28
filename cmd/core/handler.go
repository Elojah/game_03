package main

import (
	"context"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/rtc"
	"github.com/elojah/game_03/pkg/user"
	"github.com/pion/webrtc/v3"
)

type handler struct {
	entity entity.App
	room   room.App

	rtc  rtc.App
	user user.App

	api *webrtc.API
}

func (h *handler) Dial(ctx context.Context) error {
	s := webrtc.SettingEngine{}

	// s.DetachDataChannels()

	// Create an API object with the engine
	h.api = webrtc.NewAPI(webrtc.WithSettingEngine(s))

	return nil
}
