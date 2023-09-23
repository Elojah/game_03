package entity

import (
	"context"
	"math"

	"github.com/elojah/game_03/pkg/geometry"
	"github.com/elojah/game_03/pkg/ulid"
)

const (
	// MaxAbilities is a soft limit to ease caching
	MaxAbilities = 1000
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
	FactionID   ulid.ID
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

type FilterCache struct {
	ID      ulid.ID
	Min     string
	Max     string
	Size    int64
	Reverse bool
}

type Cache interface {
	InsertCache(context.Context, E) error
	FetchManyCache(context.Context, FilterCache) ([]E, error)
	DeleteCache(context.Context, FilterCache) error
}

type Store interface {
	Insert(context.Context, E) error
	Update(context.Context, Filter, Patch) error
	Fetch(context.Context, Filter) (E, error)
	FetchMany(context.Context, Filter) ([]E, []byte, error)
	Delete(context.Context, Filter) error
}

type App interface {
	Cache
	CacheAbility

	Store
	StoreAbility
	StoreAnimation
	StoreBackup
	StorePC
	StorePCPreferences
	StoreTemplate
	StoreSpawn

	CreateEntityFromBackup(context.Context, ulid.ID) (E, error)
	CreateBackupFromEntity(context.Context, ulid.ID) error

	Insert(context.Context, E) error

	CreateDefaultAbilities(context.Context, ulid.ID) ([]Ability, error)
}

// GetStat should be used when only way to get stat is `Stat` enum.
// Use direct access if you can.
func (e E) GetStat(st Stat) int64 {
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

// SetStat should be used when only way to get stat is `Stat` enum.
// Use direct access if you can.
func (e E) SetStat(st Stat, value int64) {
	switch st {
	case NoneStat:
	case Damage:
		e.Stats.Damage = value
	case Defense:
		e.Stats.Defense = value
	case MoveSpeed:
		e.Stats.MoveSpeed = value
	case CastSpeed:
		e.Stats.CastSpeed = value
	case CooldownReduction:
		e.Stats.CooldownReduction = value
	case HP:
		e.Stats.HP = value
	case MP:
		e.Stats.MP = value
	case MaxHP:
		e.Stats.MaxHP = value
	case MaxMP:
		e.Stats.MaxMP = value
	}
}

func (e E) SetPosition(x int64, y int64) {
	e.X = x
	e.Y = y
}

func (e E) Distance(comp E) float64 {
	return math.Abs(float64(comp.X-e.X)) + math.Abs(float64(comp.Y-e.Y))
}
