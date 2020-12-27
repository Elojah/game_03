package app

import "github.com/elojah/game_03/pkg/user"

var _ user.App = (*A)(nil)

type A struct {
	user.Store
	user.CacheSession
}
