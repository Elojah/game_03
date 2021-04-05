package scylla

import (
	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/user/scylla"
)

var _ entity.Store = (*Store)(nil)

var _ entity.StorePC = (*Store)(nil)

type Store struct {
	scylla.Store
}
