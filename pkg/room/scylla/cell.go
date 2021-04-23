package scylla

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/gocql/gocql"
)

type filterCell room.FilterCell

func (f filterCell) where() (string, []interface{}) {
	var clause []string

	var args []interface{}

	if len(f.Keys) > 0 {
		keys := make([]string, len(f.Keys))
		for i, k := range f.Keys {
			keys[i] = `(?, ?, ?)`

			args = append(args, k.WorldID, k.X, k.Y)
		}

		clause = append(clause, fmt.Sprintf("(world_id, x, y) IN (%s)", strings.Join(keys, ",")))
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

	if f.Keys != nil {
		ss := make([]string, 0, len(f.Keys))
		for _, k := range f.Keys {
			ss = append(ss, k.WorldID.String(), strconv.FormatInt(k.X, 10), strconv.FormatInt(k.Y, 10))
		}

		cols = append(cols, strings.Join(ss, "|"))
	}

	return strings.Join(cols, "|")
}

func (s Store) InsertCell(ctx context.Context, c room.Cell) error {
	q := s.Session.Query(
		`INSERT INTO main.cell (world_id, x, y, tilemap) VALUES (?, ?, ?, ?)`,
		c.WorldID,
		c.X,
		c.Y,
		c.Tilemap,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchCell(ctx context.Context, f room.FilterCell) (room.Cell, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT world_id, x, y, tilemap FROM main.cell `)

	clause, args := filterCell(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var c room.Cell
	if err := q.Scan(&c.WorldID, &c.X, &c.Y, &c.Tilemap); err != nil {
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
	b.WriteString(`SELECT world_id, x, y, tilemap FROM main.cell `)

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
			&cells[i].WorldID,
			&cells[i].X,
			&cells[i].Y,
			&cells[i].Tilemap,
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
