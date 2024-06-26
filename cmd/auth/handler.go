package main

import (
	"github.com/elojah/game_03/pkg/google"
	"github.com/elojah/game_03/pkg/twitch"
	"github.com/elojah/game_03/pkg/user"
)

type handler struct {
	twitch twitch.App
	google google.App

	user user.App
}
