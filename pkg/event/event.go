package event

import (
	"context"

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
