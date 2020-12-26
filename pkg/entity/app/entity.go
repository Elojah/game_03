package app

import "github.com/elojah/game_03/pkg/entity"

type App struct {
	entity.Cache
	entity.Store
	entity.StorePC
}
