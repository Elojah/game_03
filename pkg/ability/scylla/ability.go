package scylla

import (
	"context"
	"errors"
	"strings"

	"github.com/elojah/game_03/pkg/ability"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/gocql/gocql"
)

type filter ability.Filter

func (f filter) where() (string, []any) {
	var clause []string

	var args []any

	if f.ID != nil {
		clause = append(clause, `id = ?`)
		args = append(args, f.ID)
	}

	if len(f.IDs) > 0 {
		clause = append(clause, `id IN ?`)
		args = append(args, f.IDs)
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

	return strings.Join(cols, " - ")
}

func (s Store) Insert(ctx context.Context, a ability.A) error {
	raw, err := a.Marshal()
	if err != nil {
		return err
	}

	q := s.Session.Query(
		`INSERT INTO main.ability (id, proto) VALUES (?, ?)`,
		a.ID,
		raw,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) Fetch(ctx context.Context, f ability.Filter) (ability.A, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT id, proto FROM main.ability `)

	clause, args := filter(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var a ability.A

	var raw []byte
	if err := q.Scan(&a.ID, &raw); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return ability.A{}, gerrors.ErrNotFound{Resource: "tileset", Index: filter(f).index()}
		}

		return ability.A{}, err
	}

	if err := a.Unmarshal(raw); err != nil {
		return a, err
	}

	return a, nil
}

func (s Store) FetchMany(ctx context.Context, f ability.Filter) ([]ability.A, []byte, error) {
	if f.Size <= 0 {
		return nil, nil, nil
	}

	b := strings.Builder{}
	b.WriteString(`SELECT
		id, proto
	FROM main.ability `)

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

	as := make([]ability.A, f.Size)

	var i int

	for ; scanner.Next(); i++ {
		var raw []byte

		if err := scanner.Scan(
			&as[i].ID, &raw,
		); err != nil {
			return nil, nil, err
		}

		if err := as[i].Unmarshal(raw); err != nil {
			return nil, nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return as[:i], state, nil
}

func (s Store) Delete(ctx context.Context, f ability.Filter) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.ability `)

	clause, args := filter(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
