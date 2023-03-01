package scylla

import (
	"context"
	"errors"
	"strings"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/gocql/gocql"
)

type filterPublic room.FilterPublic

func (f filterPublic) where() (string, []any) {
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

	if f.RoomID != nil {
		clause = append(clause, `room_id = ?`)
		args = append(args, f.RoomID)
	}

	if len(f.RoomIDs) > 0 {
		clause = append(clause, `room_id IN ?`)
		args = append(args, f.RoomIDs)
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

func (f filterPublic) index() string {
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

	if f.RoomID != nil {
		cols = append(cols, f.RoomID.String())
	}

	if f.RoomIDs != nil {
		ss := make([]string, 0, len(f.RoomIDs))
		for _, id := range f.RoomIDs {
			ss = append(ss, id.String())
		}

		cols = append(cols, strings.Join(ss, "|"))
	}

	return strings.Join(cols, " - ")
}

func (s Store) InsertPublic(ctx context.Context, r room.Public) error {
	q := s.Session.Query(
		`INSERT INTO main.room_public (id, room_id) VALUES (?, ?)`,
		r.ID,
		r.RoomID,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchPublic(ctx context.Context, f room.FilterPublic) (room.Public, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT id, room_id FROM main.room_public `)

	clause, args := filterPublic(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var r room.Public
	if err := q.Scan(&r.ID, &r.RoomID); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return room.Public{}, gerrors.ErrNotFound{Resource: "room_public", Index: filterPublic(f).index()}
		}

		return room.Public{}, err
	}

	return r, nil
}

func (s Store) FetchManyPublic(ctx context.Context, f room.FilterPublic) ([]room.Public, []byte, error) {
	if f.Size <= 0 {
		return nil, nil, nil
	}

	b := strings.Builder{}
	b.WriteString(`SELECT id, room_id FROM main.room_public `)

	clause, args := filterPublic(f).where()
	b.WriteString(clause)

	iter := s.Session.Query(b.String(), args...).
		WithContext(ctx).
		PageState(f.State).
		PageSize(f.Size).
		Iter()

	defer iter.Close()

	state := iter.PageState()

	scanner := iter.Scanner()

	rooms := make([]room.Public, f.Size)

	var i int

	for ; scanner.Next(); i++ {
		if err := scanner.Scan(
			&rooms[i].ID,
			&rooms[i].RoomID,
		); err != nil {
			return nil, nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return rooms[:i], state, nil
}

func (s Store) DeletePublic(ctx context.Context, f room.FilterPublic) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.room_public `)

	clause, args := filterPublic(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
