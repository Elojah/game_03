package main

import (
	"github.com/elojah/game_03/pkg/cookie"
	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/migrate"
	"github.com/elojah/game_03/pkg/room"
)

type handler struct {
	migrate migrate.App

	room   room.App
	entity entity.App
	cookie cookie.App
}
