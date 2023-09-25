package scylla

import (
	"context"
	"errors"
	"strings"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/gocql/gocql"
)

type filterWorldWaypoint room.FilterWorldWaypoint

func (f filterWorldWaypoint) where() (string, []any) {
	var clause []string

	var args []any

	if f.ID != nil {
		clause = append(clause, `id = ?`)
		args = append(args, f.ID)
	}

	if f.WorldID != nil {
		clause = append(clause, `world_id = ?`)
		args = append(args, f.WorldID)
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

func (f filterWorldWaypoint) index() string {
	var cols []string

	if f.ID != nil {
		cols = append(cols, f.ID.String())
	}

	if f.WorldID != nil {
		cols = append(cols, f.WorldID.String())
	}

	return strings.Join(cols, " - ")
}

func (s Store) InsertWorldWaypoint(ctx context.Context, c room.WorldWaypoint) error {
	q := s.Session.Query(
		`INSERT INTO main.world_waypoint (id, world_id, cell_id, x, y) VALUES (?, ?, ?, ?, ?)`,
		c.ID,
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

func (s Store) FetchWorldWaypoint(ctx context.Context, f room.FilterWorldWaypoint) (room.WorldWaypoint, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT id, world_id, cell_id, x, y FROM main.world_waypoint `)

	clause, args := filterWorldWaypoint(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var c room.WorldWaypoint
	if err := q.Scan(&c.ID, &c.WorldID, &c.CellID, &c.X, &c.Y); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return room.WorldWaypoint{}, gerrors.ErrNotFound{Resource: "world_waypoint", Index: filterWorldWaypoint(f).index()}
		}

		return room.WorldWaypoint{}, err
	}

	return c, nil
}

func (s Store) FetchManyWorldWaypoint(ctx context.Context, f room.FilterWorldWaypoint) ([]room.WorldWaypoint, []byte, error) {
	if f.Size <= 0 {
		return nil, nil, nil
	}

	b := strings.Builder{}
	b.WriteString(`SELECT id, world_id, cell_id, x, y FROM main.world_waypoint `)

	clause, args := filterWorldWaypoint(f).where()
	b.WriteString(clause)

	iter := s.Session.Query(b.String(), args...).
		WithContext(ctx).
		PageState(f.State).
		PageSize(f.Size).
		Iter()

	defer iter.Close()

	state := iter.PageState()

	scanner := iter.Scanner()

	cells := make([]room.WorldWaypoint, f.Size)

	var i int

	for ; scanner.Next(); i++ {
		if err := scanner.Scan(
			&cells[i].ID,
			&cells[i].WorldID,
			&cells[i].CellID,
			&cells[i].X,
			&cells[i].Y,
		); err != nil {
			return nil, nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return cells[:i], state, nil
}

func (s Store) DeleteWorldWaypoint(ctx context.Context, f room.FilterWorldWaypoint) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.world_waypoint `)

	clause, args := filterWorldWaypoint(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
