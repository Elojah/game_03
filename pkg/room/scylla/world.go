package scylla

import (
	"context"
	"errors"
	"strings"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	gscylla "github.com/elojah/go-scylla"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/table"
)

var worldByID = gscylla.NewTable(table.Metadata{
	Name:    "world",
	Columns: []string{"id", "height", "width", "tileset"},
	PartKey: []string{"id"},
})

type filterWorld room.FilterWorld

func (f filterWorld) index() string {
	var cols []string

	if f.ID != nil {
		cols = append(cols, f.ID.String())
	}

	return strings.Join(cols, "|")
}

func (s Store) InsertWorld(ctx context.Context, w room.World) error {
	st, ns := worldByID.Ins()
	q := s.ContextQuery(ctx, st, ns).BindStruct(w)

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchWorld(ctx context.Context, f room.FilterWorld) (room.World, error) {
	st, ns := worldByID.Get()
	q := s.ContextQuery(ctx, st, ns).BindStruct(f)

	var w room.World
	if err := q.GetRelease(&w); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return room.World{}, gerrors.ErrNotFound{Resource: "world", Index: filterWorld(f).index()}
		}

		return room.World{}, err
	}

	return w, nil
}

func (s Store) DeleteWorld(ctx context.Context, f room.FilterWorld) error {
	st, ns := worldByID.Del()
	q := s.ContextQuery(ctx, st, ns).BindStruct(f)

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return nil
}
