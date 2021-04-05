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

var (
	pcByID = gscylla.NewTable(table.Metadata{
		Name:    "pc",
		Columns: []string{"id", "user_id", "room_id"},
		PartKey: []string{"id"},
	})

	pcByUserID = gscylla.NewTable(table.Metadata{
		Name:    "pc",
		Columns: []string{"id", "user_id", "room_id"},
		PartKey: []string{"user_id"},
	})

	pcByRoomID = gscylla.NewTable(table.Metadata{
		Name:    "pc",
		Columns: []string{"id", "user_id", "room_id"},
		PartKey: []string{"room_id"},
	})
)

type filterPC entity.FilterPC

func (f filterPC) table() gscylla.Table {
	if f.ID != nil && f.UserID == nil && f.RoomID == nil {
		return pcByID
	} else if f.ID == nil && f.UserID != nil && f.RoomID == nil {
		return pcByUserID
	}

	return pcByRoomID
}

func (f filterPC) index() string {
	var cols []string

	if f.ID != nil {
		cols = append(cols, f.ID.String())
	}

	return strings.Join(cols, "|")
}

func (s Store) InsertPC(ctx context.Context, pc entity.PC) error {
	st, ns := pcByID.Ins()
	q := s.ContextQuery(ctx, st, ns).BindStruct(pc)

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchPC(ctx context.Context, f entity.FilterPC) (entity.PC, error) {
	st, ns := filterPC(f).table().Get()
	q := s.ContextQuery(ctx, st, ns).BindStruct(f)

	var pc entity.PC
	if err := q.GetRelease(&pc); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return entity.PC{}, gerrors.ErrNotFound{Resource: "pc", Index: filterPC(f).index()}
		}

		return entity.PC{}, err
	}

	return pc, nil
}

func (s Store) FetchManyPC(ctx context.Context, f entity.FilterPC) ([]entity.PC, error) {
	st, ns := filterPC(f).table().Get()
	q := s.ContextQuery(ctx, st, ns).BindStruct(f)

	var pcs []entity.PC
	if err := q.SelectRelease(&pcs); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return nil, gerrors.ErrNotFound{Resource: "entity", Index: filterPC(f).index()}
		}

		return nil, err
	}

	return pcs, nil
}

func (s Store) DeletePC(ctx context.Context, f entity.FilterPC) error {
	st, ns := pcByID.Del()
	q := s.ContextQuery(ctx, st, ns).BindStruct(f)

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return nil
}
