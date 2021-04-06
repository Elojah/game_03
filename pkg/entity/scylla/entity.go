package scylla

import (
	"context"
	"errors"
	"strings"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	gscylla "github.com/elojah/go-scylla"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/table"
)

var entityByID = gscylla.NewTable(table.Metadata{
	Name:    "entity",
	Columns: []string{"id", "pc_id", "x", "y", "rot", "radius"},
	PartKey: []string{"id"},
})

type filter entity.Filter

func (f filter) index() string {
	var cols []string

	if f.ID != nil {
		cols = append(cols, f.ID.String())
	}

	return strings.Join(cols, "|")
}

func (s Store) Insert(ctx context.Context, e entity.E) error {
	st, ns := entityByID.Ins()
	q := s.ContextQuery(ctx, st, ns).BindStruct(e)

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return nil
}

func (s Store) Fetch(ctx context.Context, f entity.Filter) (entity.E, error) {
	st, ns := entityByID.Get()
	q := s.ContextQuery(ctx, st, ns).BindStruct(f)

	var e entity.E
	if err := q.GetRelease(&e); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return entity.E{}, gerrors.ErrNotFound{Resource: "entity", Index: filter(f).index()}
		}

		return entity.E{}, err
	}

	return e, nil
}

func (s Store) Delete(ctx context.Context, f entity.Filter) error {
	st, ns := entityByID.Del()
	q := s.ContextQuery(ctx, st, ns).BindStruct(f)

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return nil
}
