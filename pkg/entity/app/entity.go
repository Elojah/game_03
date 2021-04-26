package app

import "github.com/elojah/game_03/pkg/entity"

var _ entity.App = (*App)(nil)

type App struct {
	entity.Store
	entity.StorePC
	entity.CachePCConnect
}
