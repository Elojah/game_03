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

	Listen(context.Context, ulid.ID) error
}

func NewEvents(sourceID ulid.ID, c ability.Cast) map[string]E {
	result := make(map[string]E, 1)

	// #Sepcial case: abilityID = sourceID -> use this case for move events
	if target, ok := c.Targets[sourceID.String()]; ok && c.AbilityID.Compare(sourceID) == 0 {
		ev := E{
			ID:       ulid.NewID(),
			EntityID: sourceID,
			Source:   entity.E{ID: sourceID},
			At:       c.At,
			// default move effect, hard code here
			Effect: ability.CastEffect{
				ID:   ulid.NewID(),
				Self: true,
				Effect: ability.Effect{
					Targets: map[string]ability.Target{
						sourceID.String(): {
							Type: ability.Self,
							Move: ability.MoveTarget{
								Move:       ability.Walk,
								TargetType: ability.Circle,
							},
						},
					},
				},
				Targets: map[string]ability.CastTarget{
					sourceID.String(): {
						Circle: target.Circle,
					},
				},
				CurrentID: sourceID,
			},
		}

		result[ev.ID.String()] = ev

		return result
	}

	// Create source cast event for source
	ev := E{
		ID:       ulid.NewID(),
		EntityID: sourceID,
		Source:   entity.E{ID: sourceID},
		At:       c.At,
		Effect: ability.CastEffect{
			Self:      true,
			AbilityID: c.AbilityID,
		},
	}

	result[ev.ID.String()] = ev

	return result
}
