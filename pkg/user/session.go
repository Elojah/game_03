package user

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type FilterSession struct {
	ID *ulid.ID
}

type CacheSession interface {
	UpsertSession(context.Context, Session) error
	FetchSession(context.Context, FilterSession) (Session, error)
	DeleteSession(context.Context, FilterSession) error
}

type StoreSession interface {
	InsertSession(context.Context, Session) error
	FetchSession(context.Context, FilterSession) (Session, error)
	DeleteSession(context.Context, FilterSession) error
}
