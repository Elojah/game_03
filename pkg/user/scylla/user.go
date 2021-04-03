package scylla

import (
	"context"
	"fmt"

	"github.com/elojah/game_03/pkg/user"
	gscylla "github.com/elojah/go-scylla"
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

func (f filter) Table() gscylla.Table {
	if f.ID != nil && f.TwitchID == nil {
		return userByID
	} else if f.ID == nil && f.TwitchID != nil {
		return userByTwitchID
	}

	return userByIDAndTwitchID
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
	st, ns := filter(f).Table().Get()
	q := s.ContextQuery(ctx, st, ns).BindStruct(f)

	fmt.Println(q.Statement())

	var u user.U
	if err := q.GetRelease(&u); err != nil {
		return user.U{}, err
	}

	return u, nil
}

func (s Store) Delete(ctx context.Context, f user.Filter) error {
	st, ns := filter(f).Table().Get()
	q := s.ContextQuery(ctx, st, ns).BindStruct(f)

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return nil
}
