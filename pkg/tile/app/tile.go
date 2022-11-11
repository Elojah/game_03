package app

import "github.com/elojah/game_03/pkg/tile"

var _ tile.App = (*App)(nil)

type App struct {
	tile.StoreSet
}
