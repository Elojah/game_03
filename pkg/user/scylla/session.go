package scylla

import (
	"context"
	"errors"
	"strings"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/user"
	"github.com/gocql/gocql"
)

type filterSession user.FilterSession

func (f filterSession) where() (string, []any) {
	var clause []string

	var args []any

	if f.ID != nil {
		clause = append(clause, `id = ?`)
		args = append(args, f.ID)
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

func (f filterSession) index() string {
	var cols []string

	if f.ID != nil {
		cols = append(cols, f.ID.String())
	}

	return strings.Join(cols, " - ")
}

func (s Store) InsertSession(ctx context.Context, ses user.Session) error {
	q := s.Session.Query(
		`INSERT INTO main.session (id, user_id, at) VALUES (?, ?, ?)`,
		ses.ID,
		ses.UserID,
		ses.At,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchSession(ctx context.Context, f user.FilterSession) (user.Session, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT id, user_id, at FROM main.session `)

	clause, args := filterSession(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var ses user.Session
	if err := q.Scan(&ses.ID, &ses.UserID, &ses.At); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return user.Session{}, gerrors.ErrNotFound{Resource: "session", Index: filterSession(f).index()}
		}

		return user.Session{}, err
	}

	return ses, nil
}

func (s Store) DeleteSession(ctx context.Context, f user.FilterSession) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.session `)

	clause, args := filterSession(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
