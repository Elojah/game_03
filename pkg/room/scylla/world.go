package scylla

import (
	"context"
	"errors"
	"strings"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/gocql/gocql"
)

type filterWorld room.FilterWorld

func (f filterWorld) where() (string, []interface{}) {
	var clause []string

	var args []interface{}

	if f.ID != nil {
		clause = append(clause, `id = ?`)
		args = append(args, *f.ID)
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

func (f filterWorld) index() string {
	var cols []string

	if f.ID != nil {
		cols = append(cols, f.ID.String())
	}

	return strings.Join(cols, " - ")
}

func (s Store) InsertWorld(ctx context.Context, w room.World) error {
	q := s.Session.Query(
		`INSERT INTO main.world (id, height, width, cell_height, cell_width) VALUES (?, ?, ?, ?, ?)`,
		w.ID,
		w.Height,
		w.Width,
		w.CellHeight,
		w.CellWidth,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchWorld(ctx context.Context, f room.FilterWorld) (room.World, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT id, height, width, cell_height, cell_width FROM main.world `)

	clause, args := filterWorld(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var w room.World
	if err := q.Scan(&w.ID, &w.Height, &w.Width, &w.CellHeight, &w.CellWidth); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return room.World{}, gerrors.ErrNotFound{Resource: "world", Index: filterWorld(f).index()}
		}

		return room.World{}, err
	}

	return w, nil
}

func (s Store) DeleteWorld(ctx context.Context, f room.FilterWorld) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.world `)

	clause, args := filterWorld(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
