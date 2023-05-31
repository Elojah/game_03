package event

import (
	"context"

	"github.com/elojah/game_03/pkg/ability"
	entity "github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/redis/rueidis"
)

// Message alias a redis message.
type Message = rueidis.PubSubMessage

type FilterQ struct {
	EntityID ulid.ID
}

// CacheQ contains basic queue operations for event Event object.
type CacheQ interface {
	Publish(context.Context, E) error
	Subscribe(context.Context, FilterQ, func(Message)) error
}

type Filter struct {
	EntityID ulid.ID

	Min string
	Max string

	Size int64
}

// Cache contains basic operations for event Event object.
type Cache interface {
	Insert(context.Context, E) error
	FetchMany(context.Context, Filter) ([]E, error)
	Delete(context.Context, Filter) error
}

type App interface {
	Cache
	CacheQ

	Eval(context.Context, ulid.ID) error
	CreateFromCast(context.Context, ulid.ID, ability.Cast) (map[string]E, error)
}

func (ev E) Eval(e entity.E) entity.E {
	e = ev.Effect.Eval(e)

	return e
}