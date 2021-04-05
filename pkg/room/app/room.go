package app

import "github.com/elojah/game_03/pkg/room"

type App struct {
	room.Store
	room.StoreWorld
	room.StoreCell
}
