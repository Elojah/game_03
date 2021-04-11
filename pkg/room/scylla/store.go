package scylla

import (
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/go-scylla"
)

var _ room.Store = (*Store)(nil)

var _ room.StoreWorld = (*Store)(nil)

var _ room.StoreCell = (*Store)(nil)

type Store struct {
	*scylla.Service
}
