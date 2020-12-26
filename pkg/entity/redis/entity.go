package redis

import (
	"context"

	"github.com/elojah/game_03/pkg/entity"
)

const entityKey = "entity:"

func (c Cache) Upsert(context.Context, entity.E) error {
	_ = entityKey

	return nil
}

func (c Cache) Fetch(context.Context, entity.Filter) (entity.E, error) {
	return entity.E{}, nil
}

func (c Cache) FetchMany(context.Context, entity.Filter, func(entity.E) error) error {
	return nil
}

func (c Cache) Delete(context.Context, entity.Filter) error {
	return nil
}

func (c Cache) DeleteMany(context.Context, entity.Filter) error {
	return nil
}
