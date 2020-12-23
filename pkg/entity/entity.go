package entity

import (
	"context"

	"github.com/elojah/game_03/pkg/geometry"
	"github.com/elojah/game_03/pkg/ulid"
)

type Filter struct {
	ID       *ulid.ID
	KClosest *geometry.KClosest
}

type Store interface {
	Upsert(context.Context, E) error
	FetchMany(context.Context, Filter, func(E) error) error
	Delete(context.Context, Filter) error
}
