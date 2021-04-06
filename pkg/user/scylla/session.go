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

var sessionByID = gscylla.NewTable(table.Metadata{
	Name:    "session",
	Columns: []string{"id", "user_id", "twitch_token"},
	PartKey: []string{"id"},
})

type filterSession user.FilterSession

func (f filterSession) index() string {
	var cols []string

	if f.ID != nil {
		cols = append(cols, f.ID.String())
	}

	return strings.Join(cols, "|")
}

func (s Store) InsertSession(ctx context.Context, ses user.Session) error {
	st, ns := sessionByID.Ins()
	q := s.ContextQuery(ctx, st, ns).BindStruct(ses)

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchSession(ctx context.Context, f user.FilterSession) (user.Session, error) {
	st, ns := sessionByID.Get()
	q := s.ContextQuery(ctx, st, ns).BindStruct(f)

	var ses user.Session
	if err := q.GetRelease(&ses); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return user.Session{}, gerrors.ErrNotFound{Resource: "session", Index: filterSession(f).index()}
		}

		return user.Session{}, err
	}

	return ses, nil
}

func (s Store) DeleteSession(ctx context.Context, f user.FilterSession) error {
	st, ns := sessionByID.Del()
	q := s.ContextQuery(ctx, st, ns).BindStruct(f)

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return nil
}
