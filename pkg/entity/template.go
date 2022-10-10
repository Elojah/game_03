package entity

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type FilterTemplate struct {
	ID    ulid.ID
	IDs   []ulid.ID
	Name  *string
	Names []string

	State []byte
	Size  int
}

type StoreTemplate interface {
	InsertTemplate(context.Context, Template) error
	FetchTemplate(context.Context, FilterTemplate) (Template, error)
	FetchManyTemplate(context.Context, FilterTemplate) ([]Template, []byte, error)
	DeleteTemplate(context.Context, FilterTemplate) error
}
