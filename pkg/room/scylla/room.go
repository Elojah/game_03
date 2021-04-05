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
	if f.ID != nil && f.OwnerID == nil {
		return roomByID
	}

	return roomByOwnerID
}

func (f filter) index() string {
	var cols []string

	if f.ID != nil {
		cols = append(cols, f.ID.String())
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

func (s Store) FetchMany(ctx context.Context, f room.Filter) ([]room.R, error) {
	st, ns := filter(f).table().Get()
	q := s.ContextQuery(ctx, st, ns).BindStruct(f)

	var rs []room.R
	if err := q.SelectRelease(&rs); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return nil, gerrors.ErrNotFound{Resource: "room", Index: filter(f).index()}
		}

		return nil, err
	}

	return rs, nil
}

func (s Store) Delete(ctx context.Context, f room.Filter) error {
	st, ns := filter(f).table().Del()
	q := s.ContextQuery(ctx, st, ns).BindStruct(f)

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return nil
}
