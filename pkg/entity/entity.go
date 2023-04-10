package entity

import (
	"context"

	"github.com/elojah/game_03/pkg/geometry"
	"github.com/elojah/game_03/pkg/ulid"
)

type Filter struct {
	ID  ulid.ID
	IDs []ulid.ID

	CellID  ulid.ID
	CellIDs []ulid.ID

	State []byte
	Size  int
}

type Patch struct {
	UserID      ulid.ID
	CellID      ulid.ID
	Name        *string
	X           *int64
	Y           *int64
	Rot         *int32
	Radius      *int32
	At          *int64
	AnimationID ulid.ID
	AnimationAt *int64
	Objects     []geometry.Rect
}

type Store interface {
	Insert(context.Context, E) error
	Update(context.Context, Filter, Patch) error
	Fetch(context.Context, Filter) (E, error)
	FetchMany(context.Context, Filter) ([]E, []byte, error)
	Delete(context.Context, Filter) error
}

type App interface {
	Store
	StoreAnimation
	StoreBackup
	StorePC
	StoreTemplate
	StoreSpawn
}

// FetchStat should be used when only way to get stat is `Stat` enum.
// Use direct access if you can.
func (e E) FetchStat(st Stat) int64 {
	switch st {
	case NoneStat:
		return 0
	case Damage:
		return e.Stats.Damage
	case Defense:
		return e.Stats.Defense
	case MoveSpeed:
		return e.Stats.MoveSpeed
	case CastSpeed:
		return e.Stats.CastSpeed
	case CooldownReduction:
		return e.Stats.CooldownReduction
	case HP:
		return e.Stats.HP
	case MP:
		return e.Stats.MP
	case MaxHP:
		return e.Stats.MaxHP
	case MaxMP:
		return e.Stats.MaxMP
	}

	return 0
}
