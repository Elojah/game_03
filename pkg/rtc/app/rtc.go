package app

import "github.com/elojah/game_03/pkg/rtc"

var _ rtc.App = (*App)(nil)

type App struct {
	rtc.Store
}
