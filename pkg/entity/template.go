package entity

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type FilterTemplate struct {
	EntityID  ulid.ID
	EntityIDs []ulid.ID
	Name      *string
	Names     []string

	All bool

	State []byte
	Size  int
}

type StoreTemplate interface {
	InsertTemplate(context.Context, Template) error
	FetchTemplate(context.Context, FilterTemplate) (Template, error)
	FetchManyTemplate(context.Context, FilterTemplate) ([]Template, []byte, error)
	DeleteTemplate(context.Context, FilterTemplate) error
}
