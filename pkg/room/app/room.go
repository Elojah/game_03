package app

import (
	"github.com/elojah/game_03/pkg/room"
)

var _ room.App = (*App)(nil)

type App struct {
	room.Store
	room.StoreWorld
	room.StoreCell
}
