package main

import (
	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/space"
	"github.com/elojah/game_03/pkg/twitch"
	"github.com/elojah/game_03/pkg/user"
)

type handler struct {
	entity entity.App // nolint: structcheck
	space  space.App  // nolint: structcheck
	twitch twitch.App
	user   user.App
}
