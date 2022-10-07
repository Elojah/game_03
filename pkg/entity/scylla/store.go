package scylla

import (
	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/go-scylla"
)

var _ entity.Store = (*Store)(nil)

var _ entity.StoreAnimation = (*Store)(nil)

var _ entity.StoreBackup = (*Store)(nil)

var _ entity.StorePC = (*Store)(nil)

var _ entity.StoreNPC = (*Store)(nil)

type Store struct {
	scylla.Service
}
