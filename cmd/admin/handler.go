package main

import (
	"github.com/elojah/game_03/pkg/ability"
	"github.com/elojah/game_03/pkg/cookie"
	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/migrate"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/tile"
)

type handler struct {
	migrate migrate.App

	ability ability.App
	cookie  cookie.App
	entity  entity.App
	room    room.App
	tile    tile.App
}
