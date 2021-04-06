package scylla

import (
	"github.com/elojah/game_03/pkg/user"
	"github.com/elojah/go-scylla"
)

var _ user.Store = (*Store)(nil)

type Store struct {
	*scylla.Service
}
