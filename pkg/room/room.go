package room

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type Filter struct {
	ID       ulid.ID
	IDs      []ulid.ID
	OwnerID  ulid.ID
	OwnerIDs []ulid.ID

	All bool

	State []byte
	Size  int
}

type Store interface {
	Insert(context.Context, R) error
	Fetch(context.Context, Filter) (R, error)
	FetchMany(context.Context, Filter) ([]R, []byte, error)
	Delete(context.Context, Filter) error
}

type FilterPublic struct {
	ID      ulid.ID
	IDs     []ulid.ID
	RoomID  ulid.ID
	RoomIDs []ulid.ID

	All bool

	State []byte
	Size  int
}

type StorePublic interface {
	InsertPublic(context.Context, Public) error
	FetchPublic(context.Context, FilterPublic) (Public, error)
	FetchManyPublic(context.Context, FilterPublic) ([]Public, []byte, error)
	DeletePublic(context.Context, FilterPublic) error
}

type App interface {
	Store
	StorePublic
	StoreCell
	StoreUser
	StoreWorld
	StoreWorldCell
	StoreWorldSpawn

	PopulateWaypoints(context.Context, Waypoints, ulid.ID) error
	CopyWorld(context.Context, ulid.ID) (ulid.ID, error)
}
