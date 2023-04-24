package main

import (
	"github.com/elojah/game_03/pkg/rtc"
	"github.com/elojah/game_03/pkg/user"
)

type handler struct {
	user user.App
	rtc  rtc.App
}
