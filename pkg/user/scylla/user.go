package scylla

import (
	"context"
	"errors"
	"strings"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/user"
	"github.com/gocql/gocql"
)

type filter user.Filter

func (f filter) where() (string, []interface{}) {
	var clause []string

	var args []interface{}

	if f.ID != nil {
		clause = append(clause, `id = ?`)
		args = append(args, *f.ID)
	}

	if len(f.IDs) > 0 {
		clause = append(clause, `id IN ?`)
		args = append(args, f.IDs)
	}

	if f.TwitchID != nil {
		clause = append(clause, `twitch_id = ?`)
		args = append(args, *f.TwitchID)
	}

	b := strings.Builder{}
	b.WriteString(" WHERE ")

	if len(clause) == 0 {
		b.WriteString("false")
	} else {
		b.WriteString(strings.Join(clause, " AND "))
	}

	return b.String(), args
}

func (f filter) index() string {
	var cols []string

	if f.ID != nil {
		cols = append(cols, f.ID.String())
	}

	if f.IDs != nil {
		ss := make([]string, 0, len(f.IDs))
		for _, id := range f.IDs {
			ss = append(ss, id.String())
		}

		cols = append(cols, strings.Join(ss, "|"))
	}

	if f.TwitchID != nil {
		cols = append(cols, *f.TwitchID)
	}

	return strings.Join(cols, " - ")
}

func (s Store) Insert(ctx context.Context, u user.U) error {
	q := s.Session.Query(
		`INSERT INTO main.user (id, twitch_id) VALUES (?, ?)`,
		u.ID,
		u.TwitchID,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) Fetch(ctx context.Context, f user.Filter) (user.U, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT id, twitch_id FROM main.user `)

	clause, args := filter(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var u user.U
	if err := q.Scan(&u.ID, &u.TwitchID); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return user.U{}, gerrors.ErrNotFound{Resource: "user", Index: filter(f).index()}
		}

		return user.U{}, err
	}

	return u, nil
}

func (s Store) FetchMany(ctx context.Context, f user.Filter) ([]user.U, []byte, error) {
	if f.Size == 0 {
		return nil, nil, nil
	}

	b := strings.Builder{}
	b.WriteString(`SELECT id, twitch_id FROM main.user `)

	clause, args := filter(f).where()
	b.WriteString(clause)

	iter := s.Session.Query(b.String(), args...).
		WithContext(ctx).
		PageState(f.State).
		PageSize(f.Size).
		Iter()

	defer iter.Close()

	state := iter.PageState()

	scanner := iter.Scanner()

	users := make([]user.U, f.Size)

	var i int

	for ; scanner.Next(); i++ {
		if err := scanner.Scan(
			&users[i].ID,
			&users[i].TwitchID,
		); err != nil {
			return nil, nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return users[:i], state, nil
}

func (s Store) Delete(ctx context.Context, f user.Filter) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.user `)

	clause, args := filter(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
