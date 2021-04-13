package scylla

import (
	"context"
	"errors"
	"strconv"
	"strings"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	gscylla "github.com/elojah/go-scylla"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/table"
)

var cellByID = gscylla.NewTable(table.Metadata{
	Name:    "cell",
	Columns: []string{"id", "height", "width", "tileset"},
	PartKey: []string{"world_id", "x", "y"},
})

type filterCell room.FilterCell

func (f filterCell) index() string {
	var cols []string

	if f.WorldID != nil {
		cols = append(cols, f.WorldID.String())
	}

	if f.X != nil {
		cols = append(cols, strconv.FormatInt(*f.X, 10))
	}

	if f.Y != nil {
		cols = append(cols, strconv.FormatInt(*f.Y, 10))
	}

	return strings.Join(cols, "|")
}

func (s Store) InsertCells(ctx context.Context, cs ...room.Cell) error {

	st, ns := cellByID.Ins()
	b := s.Session.NewBatch(gocql.UnloggedBatch).WithContext(ctx)
	b.Bind(st, func(q *gocql.QueryInfo) ([]interface{}, error) {
		// q.
	})

	q := s.ContextQuery(ctx, st, ns).BindStruct(cs)

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchCell(ctx context.Context, f room.FilterCell) (room.Cell, error) {
	st, ns := cellByID.Get()
	q := s.ContextQuery(ctx, st, ns).BindStruct(f)

	var c room.Cell
	if err := q.GetRelease(&c); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return room.Cell{}, gerrors.ErrNotFound{Resource: "cell", Index: filterCell(f).index()}
		}

		return room.Cell{}, err
	}

	return c, nil
}

func (s Store) DeleteCell(ctx context.Context, f room.FilterCell) error {
	st, ns := cellByID.Del()
	q := s.ContextQuery(ctx, st, ns).BindStruct(f)

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return nil
}
