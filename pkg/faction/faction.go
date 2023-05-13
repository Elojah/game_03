package faction

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type Permission int64

const (
	NonePermission Permission = iota
	Guest
	Member
	Officer
	Admin
	Owner
)

type Filter struct {
	ID  ulid.ID
	IDs []ulid.ID

	WorldID  ulid.ID
	WorldIDs []ulid.ID

	State []byte
	Size  int
}

type Patch struct {
	Name *string
	Icon *string
}

type Store interface {
	Insert(context.Context, F) error
	Update(context.Context, Filter, Patch) error
	Fetch(context.Context, Filter) (F, error)
	FetchMany(context.Context, Filter) ([]F, []byte, error)
	Delete(context.Context, Filter) error
}

type FilterPC struct {
	ID  ulid.ID
	IDs []ulid.ID

	FactionID  ulid.ID
	FactionIDs []ulid.ID

	State []byte
	Size  int
}

type PatchPC struct {
	FactionID  ulid.ID
	Permission *string
}

type StorePC interface {
	InsertPC(context.Context, PC) error
	UpdatePC(context.Context, FilterPC, PatchPC) error
	FetchPC(context.Context, FilterPC) (PC, error)
	FetchManyPC(context.Context, FilterPC) ([]PC, []byte, error)
	DeletePC(context.Context, FilterPC) error
}

type App interface {
	Store
	StorePC
}
