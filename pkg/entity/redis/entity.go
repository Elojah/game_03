package redis

import (
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/ulid"
)

const (
	key = "entity:"
)

// Upsert implemented with redis.
func (s *Store) Upsert(ctx context.Context, e entity.E) error {
	raw, err := e.Marshal()
	if err != nil {
		return fmt.Errorf("upsert entity %s at %d: %w", e.ID.String(), e.At, err)
	}

	if err := s.ZAddNX(
		ctx,
		key+e.ID.String(),
		redis.Z{
			Score:  float64(e.At),
			Member: raw,
		},
	).Err(); err != nil {
		return fmt.Errorf("upsert entity %s at %d: %w", e.ID.String(), e.At, err)
	}

	return nil
}

// Fetch implemented with redis.
func (s *Store) Fetch(id ulid.ID, entityID ulid.ID) (entity.E, error) {
	vals, err := s.ZRangeByScore(
		key+entityID.String(),
		redis.ZRangeBy{
			Min:   strconv.FormatUint(id.Time(), 10),
			Max:   "+inf",
			Count: 1,
		},
	).Result()
	if err != nil {
		return entity.E{}, fmt.Errorf("fetch entity for entity %s at %d", entityID.String(), id.Time(), err)
	}
	if len(vals) != 0 {
		return entity.E{}, fmt.Errorf(gerrors.ErrNotFound{
			Store: key,
			Index: key + entityID.String(),
		}, "fetch entity at %d", id.Time())
	}

	var e entity.E
	if err := e.Unmarshal([]byte(vals[0])); err != nil {
		return entity.E{}, fmt.Errorf("fetch entity for entity %s at %d", entityID.String(), id.Time(), err)
	}
	return e, nil
}

// List list events in redis upsert key from min (included).
func (s *Store) List(id ulid.ID, min ulid.ID) ([]entity.E, error) {
	vals, err := s.ZRangeByScore(
		key+id.String(),
		redis.ZRangeBy{
			Min: strconv.FormatUint(min.Time(), 10),
			Max: "+inf",
		},
	).Result()
	if err != nil {
		return nil, fmt.Errorf("list events for entity %s at %d", id.String(), min.Time(), err)
	}
	events := make([]entity.E, len(vals))
	for i := range vals {
		if err := events[i].Unmarshal([]byte(vals[i])); err != nil {
			return nil, fmt.Errorf("list events for entity %s at %d", id.String(), min.Time(), err)
		}
	}
	return events, nil
}

// Remove implemented with redis.
func (s *Store) Remove(eventID ulid.ID, id ulid.ID) error {
	return fmt.Errorf(s.ZRem(
		key+id.String(),
		eventID.String(),
	).Err(), "remove entity %s for entity %s", eventID.String(), id.String())
}
