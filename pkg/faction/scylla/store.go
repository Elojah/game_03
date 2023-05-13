package scylla

import (
	"github.com/elojah/game_03/pkg/faction"
	"github.com/elojah/go-scylla"
)

var _ faction.Store = (*Store)(nil)

var _ faction.StorePC = (*Store)(nil)

type Store struct {
	scylla.Service
}
