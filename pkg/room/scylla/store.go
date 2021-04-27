package scylla

import (
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/go-scylla"
)

var _ room.Store = (*Store)(nil)

var _ room.StoreCell = (*Store)(nil)

var _ room.StoreUser = (*Store)(nil)

var _ room.StoreWorld = (*Store)(nil)

var _ room.StoreWorldCell = (*Store)(nil)

type Store struct {
	scylla.Service
}
