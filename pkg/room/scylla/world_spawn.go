package scylla

import (
	"context"
	"errors"
	"strings"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/gocql/gocql"
)

type filterWorldSpawn room.FilterWorldSpawn

func (f filterWorldSpawn) where() (string, []any) {
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

	if f.WorldID != nil {
		clause = append(clause, `world_id = ?`)
		args = append(args, f.WorldID)
	}

	if len(f.WorldIDs) > 0 {
		clause = append(clause, `world_id IN ?`)
		args = append(args, f.WorldIDs)
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

func (f filterWorldSpawn) index() string {
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

	if f.WorldID != nil {
		cols = append(cols, f.WorldID.String())
	}

	if f.WorldIDs != nil {
		ss := make([]string, 0, len(f.WorldIDs))
		for _, id := range f.WorldIDs {
			ss = append(ss, id.String())
		}

		cols = append(cols, strings.Join(ss, "|"))
	}

	return strings.Join(cols, " - ")
}

func (s Store) InsertWorldSpawn(ctx context.Context, c room.WorldSpawn) error {
	q := s.Session.Query(
		`INSERT INTO main.world_spawn (id, world_id, entity_id, owner_id) VALUES (?, ?, ?, ?)`,
		c.ID,
		c.WorldID,
		c.EntityID,
		c.OwnerID,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchWorldSpawn(ctx context.Context, f room.FilterWorldSpawn) (room.WorldSpawn, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT id, world_id, entity_id, owner_id FROM main.world_spawn `)

	clause, args := filterWorldSpawn(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var c room.WorldSpawn
	if err := q.Scan(&c.ID, &c.WorldID, &c.EntityID, &c.OwnerID); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return room.WorldSpawn{}, gerrors.ErrNotFound{Resource: "world_spawn", Index: filterWorldSpawn(f).index()}
		}

		return room.WorldSpawn{}, err
	}

	return c, nil
}

func (s Store) FetchManyWorldSpawn(ctx context.Context, f room.FilterWorldSpawn) ([]room.WorldSpawn, []byte, error) {
	if f.Size <= 0 {
		return nil, nil, nil
	}

	b := strings.Builder{}
	b.WriteString(`SELECT id, world_id, entity_id, owner_id FROM main.world_spawn `)

	clause, args := filterWorldSpawn(f).where()
	b.WriteString(clause)

	iter := s.Session.Query(b.String(), args...).
		WithContext(ctx).
		PageState(f.State).
		PageSize(f.Size).
		Iter()

	defer iter.Close()

	state := iter.PageState()

	scanner := iter.Scanner()

	cells := make([]room.WorldSpawn, f.Size)

	var i int

	for ; scanner.Next(); i++ {
		if err := scanner.Scan(
			&cells[i].ID,
			&cells[i].WorldID,
			&cells[i].EntityID,
			&cells[i].OwnerID,
		); err != nil {
			return nil, nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return cells[:i], state, nil
}

func (s Store) DeleteWorldSpawn(ctx context.Context, f room.FilterWorldSpawn) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.world_spawn `)

	clause, args := filterWorldSpawn(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
