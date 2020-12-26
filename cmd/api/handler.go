package main

import (
	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/space"
	"github.com/elojah/game_03/pkg/user"
)

type handler struct {
	user   user.App
	entity entity.App // nolint: structcheck
	space  space.App  // nolint: structcheck
}
