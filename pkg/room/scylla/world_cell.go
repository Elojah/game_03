package scylla

import (
	"context"
	"errors"
	"strconv"
	"strings"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/gocql/gocql"
)

type filterWorldCell room.FilterWorldCell

func (f filterWorldCell) where() (string, []interface{}) {
	var clause []string

	var args []interface{}

	clause = append(clause, `world_id = ?`)
	args = append(args, f.WorldID)

	clause = append(clause, `x = ?`)
	args = append(args, f.X)

	clause = append(clause, `y = ?`)
	args = append(args, f.Y)

	b := strings.Builder{}
	b.WriteString(" WHERE ")

	if len(clause) == 0 {
		b.WriteString("false")
	} else {
		b.WriteString(strings.Join(clause, " AND "))
	}

	return b.String(), args
}

func (f filterWorldCell) index() string {
	var cols []string

	cols = append(cols, f.WorldID.String())
	cols = append(cols, strconv.FormatInt(f.X, 10))
	cols = append(cols, strconv.FormatInt(f.Y, 10))

	return strings.Join(cols, "|")
}

func (s Store) InsertWorldCell(ctx context.Context, c room.WorldCell) error {
	q := s.Session.Query(
		`INSERT INTO main.world_cell (world_id, cell_id, x, y) VALUES (?, ?, ?, ?)`,
		c.WorldID,
		c.CellID,
		c.X,
		c.Y,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchWorldCell(ctx context.Context, f room.FilterWorldCell) (room.WorldCell, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT world_id, cell_id, x, y FROM main.world_cell `)

	clause, args := filterWorldCell(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var c room.WorldCell
	if err := q.Scan(&c.WorldID, &c.CellID, &c.X, &c.Y); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return room.WorldCell{}, gerrors.ErrNotFound{Resource: "world_cell", Index: filterWorldCell(f).index()}
		}

		return room.WorldCell{}, err
	}

	return c, nil
}

func (s Store) DeleteWorldCell(ctx context.Context, f room.FilterWorldCell) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.world_cell `)

	clause, args := filterWorldCell(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
