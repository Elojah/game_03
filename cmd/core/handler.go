package main

import (
	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/rtc"
	"github.com/elojah/game_03/pkg/user"
)

type handler struct {
	entity entity.App
	rtc    rtc.App
	user   user.App
}
