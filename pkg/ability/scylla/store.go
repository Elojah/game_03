package scylla

import (
	"github.com/elojah/game_03/pkg/ability"
	"github.com/elojah/go-scylla"
)

var _ ability.Store = (*Store)(nil)

type Store struct {
	scylla.Service
}
