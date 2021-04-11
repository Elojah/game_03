package main

import (
	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/twitch"
	"github.com/elojah/game_03/pkg/user"
)

type handler struct {
	twitch twitch.App
	user   user.App
	room   room.App
	entity entity.App
}
