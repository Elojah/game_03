package scylla

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/tile"
	"github.com/gocql/gocql"
)

type filter tile.FilterSet

func (f filter) where() (string, []any) {
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

func (f filter) index() string {
	var cols []string

	if f.ID != nil {
		cols = append(cols, f.ID.String())
	}

	return strings.Join(cols, " - ")
}

func (s Store) InsertSet(ctx context.Context, set tile.Set) error {
	raw, err := json.Marshal(set)
	if err != nil {
		return err
	}

	q := s.Session.Query(
		`INSERT INTO main.tileset (id, json) VALUES (?, ?)`,
		set.ID,
		string(raw),
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchSet(ctx context.Context, f tile.FilterSet) (tile.Set, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT id, json FROM main.tileset `)

	clause, args := filter(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var set tile.Set

	var raw string
	if err := q.Scan(&set.ID, &raw); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return tile.Set{}, gerrors.ErrNotFound{Resource: "tileset", Index: filter(f).index()}
		}

		return tile.Set{}, err
	}

	if err := json.Unmarshal([]byte(raw), &set); err != nil {
		return set, err
	}

	return set, nil
}

func (s Store) DeleteSet(ctx context.Context, f tile.FilterSet) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.tileset `)

	clause, args := filter(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
