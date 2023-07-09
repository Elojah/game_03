package ability

import (
	"context"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/ulid"
)

var (
	tmpFetchEntity         = func(ulid.ID) (entity.E, error) { return entity.E{}, nil }
	tmpUpdateEntity        = func(entity.E) error { return nil }
	tmpFetchEntityAbility  = func(ulid.ID, ulid.ID) (entity.Ability, error) { return entity.Ability{}, nil }
	tmpUpdateEntityAbility = func(entity.Ability) error { return nil }

	operators = map[Operator]func(a int64, b int64) bool{
		NoneOperator: func(a, b int64) bool { return false },
		Equal:        func(a, b int64) bool { return a == b },
		NotEqual:     func(a, b int64) bool { return a != b },
		Greater:      func(a, b int64) bool { return a > b },
		Lesser:       func(a, b int64) bool { return a < b },
	}
)

type Filter struct {
	ID  ulid.ID
	IDs []ulid.ID

	State []byte
	Size  int
}

type Store interface {
	Insert(context.Context, A) error
	Fetch(context.Context, Filter) (A, error)
	FetchMany(context.Context, Filter) ([]A, []byte, error)
	Delete(context.Context, Filter) error
}

// TODO
type Cache interface{}

type App interface {
	Cache
	Store

	AddCast(context.Context, entity.E, Cast) (entity.E, error)
}
