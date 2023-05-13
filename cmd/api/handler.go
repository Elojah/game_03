package main

import (
	"github.com/elojah/game_03/pkg/ability"
	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/faction"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/user"
)

type handler struct {
	ability ability.App
	entity  entity.App
	faction faction.App
	room    room.App
	user    user.App
}
