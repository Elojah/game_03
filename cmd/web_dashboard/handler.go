package main

import (
	"github.com/elojah/game_03/cmd/auth/grpc"
	"github.com/elojah/game_03/pkg/cookie"
	"github.com/elojah/game_03/pkg/google"
	"github.com/elojah/game_03/pkg/twitch"
)

type handler struct {
	cookie cookie.App
	twitch twitch.App
	google google.App

	grpc.AuthClient
}
