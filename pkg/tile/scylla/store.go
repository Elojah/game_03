package scylla

import (
	"github.com/elojah/game_03/pkg/tile"
	"github.com/elojah/go-scylla"
)

var _ tile.StoreSet = (*Store)(nil)

type Store struct {
	scylla.Service
}
