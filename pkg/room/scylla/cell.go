package scylla

import (
	"context"
	"errors"
	"strings"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/gocql/gocql"
)

type filterCell room.FilterCell

func (f filterCell) where() (string, []interface{}) {
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

	b := strings.Builder{}
	b.WriteString(" WHERE ")

	if len(clause) == 0 {
		b.WriteString("false")
	} else {
		b.WriteString(strings.Join(clause, " AND "))
	}

	return b.String(), args
}

func (f filterCell) index() string {
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

func (s Store) InsertCell(ctx context.Context, c room.Cell) error {
	q := s.Session.Query(
		`INSERT INTO main.cell (id, contiguous, tilemap, tileset) VALUES (?, ?, ?, ?)`,
		c.ID,
		c.Contiguous,
		c.Tilemap,
		c.Tileset,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchCell(ctx context.Context, f room.FilterCell) (room.Cell, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT id, contiguous, tilemap, tileset FROM main.cell `)

	clause, args := filterCell(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var c room.Cell
	if err := q.Scan(&c.ID, &c.Contiguous, &c.Tilemap, &c.Tileset); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return room.Cell{}, gerrors.ErrNotFound{Resource: "cell", Index: filterCell(f).index()}
		}

		return room.Cell{}, err
	}

	return c, nil
}

func (s Store) FetchManyCell(ctx context.Context, f room.FilterCell) ([]room.Cell, []byte, error) {
	if f.Size == 0 {
		return nil, nil, nil
	}

	b := strings.Builder{}
	b.WriteString(`SELECT id, contiguous, tilemap, tileset FROM main.cell `)

	clause, args := filterCell(f).where()
	b.WriteString(clause)

	iter := s.Session.Query(b.String(), args...).
		WithContext(ctx).
		PageState(f.State).
		PageSize(f.Size).
		Iter()

	defer iter.Close()

	state := iter.PageState()

	scanner := iter.Scanner()

	cells := make([]room.Cell, f.Size)

	var i int

	for ; scanner.Next(); i++ {
		if err := scanner.Scan(
			&cells[i].ID,
			&cells[i].Contiguous,
			&cells[i].Tilemap,
			&cells[i].Tileset,
		); err != nil {
			return nil, nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return cells[:i], state, nil
}

func (s Store) DeleteCell(ctx context.Context, f room.FilterCell) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.cell `)

	clause, args := filterCell(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
