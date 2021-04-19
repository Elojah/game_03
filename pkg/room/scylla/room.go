package scylla

import (
	"context"
	"errors"
	"fmt"
	"strings"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	gscylla "github.com/elojah/go-scylla"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/table"
)

var (
	roomByID = gscylla.NewTable(table.Metadata{
		Name:    "room",
		Columns: []string{"id", "owner_id", "world_id"},
		PartKey: []string{"id"},
	})

	roomByOwnerID = gscylla.NewTable(table.Metadata{
		Name:    "room",
		Columns: []string{"id", "owner_id", "world_id"},
		PartKey: []string{"owner_id"},
	})
)

type filter room.Filter

func (f filter) table() gscylla.Table {
	if len(f.IDs) > 0 {
		return roomByID
	}

	if f.OwnerID != nil {
		return roomByOwnerID
	}

	return roomByID
}

func (f filter) index() string {
	var cols []string

	if f.IDs != nil {
		ss := make([]string, 0, len(f.IDs))
		for _, id := range f.IDs {
			ss = append(ss, id.String())
		}

		cols = append(cols, strings.Join(ss, "|"))
	}

	if f.OwnerID != nil {
		cols = append(cols, f.OwnerID.String())
	}

	return strings.Join(cols, "|")
}

func (s Store) Insert(ctx context.Context, r room.R) error {
	st, ns := roomByID.Ins()
	q := s.ContextQuery(ctx, st, ns).BindStruct(r)

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return nil
}

func (s Store) Fetch(ctx context.Context, f room.Filter) (room.R, error) {
	st, ns := filter(f).table().Get()
	q := s.ContextQuery(ctx, st, ns).BindStruct(f)

	var r room.R
	if err := q.GetRelease(&r); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return room.R{}, gerrors.ErrNotFound{Resource: "room", Index: filter(f).index()}
		}

		return room.R{}, err
	}

	return r, nil
}

func (s Store) FetchMany(ctx context.Context, f room.Filter) ([]room.R, []byte, error) {
	st, ns := filter(f).table().Get()
	q := s.ContextQuery(ctx, st, ns).
		// PageSize(f.Size).
		// PageState(f.State).
		BindStruct(f)

	cursor := q.Iter().PageState()

	fmt.Println(q.String())

	var rs []room.R
	if err := q.SelectRelease(&rs); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return nil, nil, gerrors.ErrNotFound{Resource: "room", Index: filter(f).index()}
		}

		return nil, nil, err
	}

	return rs, cursor, nil
}

func (s Store) Delete(ctx context.Context, f room.Filter) error {
	st, ns := filter(f).table().Del()
	q := s.ContextQuery(ctx, st, ns).BindStruct(f)

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return nil
}
