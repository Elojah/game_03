package scylla

import (
	"context"
	"errors"
	"strings"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/user"
	gscylla "github.com/elojah/go-scylla"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/table"
)

var (
	userByID = gscylla.NewTable(table.Metadata{
		Name:    "user",
		Columns: []string{"id", "twitch_id"},
		PartKey: []string{"id"},
	})

	userByTwitchID = gscylla.NewTable(table.Metadata{
		Name:    "user",
		Columns: []string{"id", "twitch_id"},
		PartKey: []string{"twitch_id"},
	})

	userByIDAndTwitchID = gscylla.NewTable(table.Metadata{
		Name:    "user",
		Columns: []string{"id", "twitch_id"},
		PartKey: []string{"id", "twitch_id"},
	})
)

type filter user.Filter

func (f filter) table() gscylla.Table {
	if f.ID != nil && f.TwitchID == nil {
		return userByID
	} else if f.ID == nil && f.TwitchID != nil {
		return userByTwitchID
	}

	return userByIDAndTwitchID
}

func (f filter) index() string {
	var cols []string

	if f.ID != nil {
		cols = append(cols, f.ID.String())
	}

	if f.TwitchID != nil {
		cols = append(cols, *f.TwitchID)
	}

	return strings.Join(cols, "|")
}

func (s Store) Insert(ctx context.Context, u user.U) error {
	st, ns := userByID.Ins()
	q := s.ContextQuery(ctx, st, ns).BindStruct(u)

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return nil
}

func (s Store) Fetch(ctx context.Context, f user.Filter) (user.U, error) {
	st, ns := filter(f).table().Get()
	q := s.ContextQuery(ctx, st, ns).BindStruct(f)

	var u user.U
	if err := q.GetRelease(&u); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return user.U{}, gerrors.ErrNotFound{Resource: "user", Index: filter(f).index()}
		}

		return user.U{}, err
	}

	return u, nil
}

func (s Store) FetchMany(ctx context.Context, f user.Filter) ([]user.U, error) {
	st, ns := filter(f).table().Get()
	q := s.ContextQuery(ctx, st, ns).BindStruct(f)

	var us []user.U
	if err := q.SelectRelease(&us); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return nil, gerrors.ErrNotFound{Resource: "user", Index: filter(f).index()}
		}

		return nil, err
	}

	return us, nil
}

func (s Store) Delete(ctx context.Context, f user.Filter) error {
	st, ns := filter(f).table().Del()
	q := s.ContextQuery(ctx, st, ns).BindStruct(f)

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return nil
}
