package app

import "github.com/elojah/game_03/pkg/faction"

var _ faction.App = (*App)(nil)

type App struct {
	faction.Store
	faction.StorePC
}
