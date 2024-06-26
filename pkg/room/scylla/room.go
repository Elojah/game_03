package scylla

import (
	"context"
	"errors"
	"strings"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/gocql/gocql"
)

type filter room.Filter

func (f filter) where() (string, []any) {
	var clause []string

	var args []any

	if f.All {
		return "", nil
	}

	if f.ID != nil {
		clause = append(clause, `id = ?`)
		args = append(args, f.ID)
	}

	if len(f.IDs) > 0 {
		clause = append(clause, `id IN ?`)
		args = append(args, f.IDs)
	}

	if f.OwnerID != nil {
		clause = append(clause, `owner_id = ?`)
		args = append(args, f.OwnerID)
	}

	if len(f.OwnerIDs) > 0 {
		clause = append(clause, `owner_id IN ?`)
		args = append(args, f.OwnerIDs)
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

	if f.OwnerID != nil {
		cols = append(cols, f.OwnerID.String())
	}

	if f.OwnerIDs != nil {
		ss := make([]string, 0, len(f.OwnerIDs))
		for _, id := range f.OwnerIDs {
			ss = append(ss, id.String())
		}

		cols = append(cols, strings.Join(ss, "|"))
	}

	return strings.Join(cols, " - ")
}

func (s Store) Insert(ctx context.Context, r room.R) error {
	q := s.Session.Query(
		`INSERT INTO main.room (id, owner_id, world_id, name) VALUES (?, ?, ?, ?)`,
		r.ID,
		r.OwnerID,
		r.WorldID,
		r.Name,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) Fetch(ctx context.Context, f room.Filter) (room.R, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT id, owner_id, world_id, name FROM main.room `)

	clause, args := filter(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var r room.R
	if err := q.Scan(&r.ID, &r.OwnerID, &r.WorldID, &r.Name); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return room.R{}, gerrors.ErrNotFound{Resource: "room", Index: filter(f).index()}
		}

		return room.R{}, err
	}

	return r, nil
}

func (s Store) FetchMany(ctx context.Context, f room.Filter) ([]room.R, []byte, error) {
	if f.Size <= 0 {
		return nil, nil, nil
	}

	b := strings.Builder{}
	b.WriteString(`SELECT id, owner_id, world_id, name FROM main.room `)

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

	rooms := make([]room.R, f.Size)

	var i int

	for ; scanner.Next(); i++ {
		if err := scanner.Scan(
			&rooms[i].ID,
			&rooms[i].OwnerID,
			&rooms[i].WorldID,
			&rooms[i].Name,
		); err != nil {
			return nil, nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return rooms[:i], state, nil
}

func (s Store) Delete(ctx context.Context, f room.Filter) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.room `)

	clause, args := filter(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
