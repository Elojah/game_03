package scylla

import (
	"context"
	"errors"
	"strings"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/gocql/gocql"
)

type filterPC entity.FilterPC

func (f filterPC) where() (string, []interface{}) {
	var clause []string

	var args []interface{}

	if f.ID != nil {
		clause = append(clause, `id = ?`)
		args = append(args, f.ID)
	}

	if len(f.IDs) > 0 {
		clause = append(clause, `id IN ?`)
		args = append(args, f.IDs)
	}

	if f.UserID != nil {
		clause = append(clause, `user_id = ?`)
		args = append(args, f.UserID)
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

func (f filterPC) index() string {
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

	if f.UserID != nil {
		cols = append(cols, f.UserID.String())
	}

	if f.WorldID != nil {
		cols = append(cols, f.WorldID.String())
	}

	return strings.Join(cols, " - ")
}

func (s Store) InsertPC(ctx context.Context, pc entity.PC) error {
	q := s.Session.Query(
		`INSERT INTO main.pc (id, user_id, world_id, entity_id) VALUES (?, ?, ?, ?)`,
		pc.ID,
		pc.UserID,
		pc.WorldID,
		pc.EntityID,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchPC(ctx context.Context, f entity.FilterPC) (entity.PC, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT id, user_id, world_id, entity_id FROM main.pc `)

	clause, args := filterPC(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var pc entity.PC
	if err := q.Scan(&pc.ID, &pc.UserID, &pc.WorldID, &pc.EntityID); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return entity.PC{}, gerrors.ErrNotFound{Resource: "pc", Index: filterPC(f).index()}
		}

		return entity.PC{}, err
	}

	return pc, nil
}

func (s Store) FetchManyPC(ctx context.Context, f entity.FilterPC) ([]entity.PC, []byte, error) {
	if f.Size <= 0 {
		return nil, nil, nil
	}

	b := strings.Builder{}
	b.WriteString(`SELECT id, user_id, world_id, entity_id FROM main.pc `)

	clause, args := filterPC(f).where()
	b.WriteString(clause)

	iter := s.Session.Query(b.String(), args...).
		WithContext(ctx).
		PageState(f.State).
		PageSize(f.Size).
		Iter()

	defer iter.Close()

	state := iter.PageState()

	scanner := iter.Scanner()

	pcs := make([]entity.PC, f.Size)

	var i int

	for ; scanner.Next(); i++ {
		if err := scanner.Scan(
			&pcs[i].ID,
			&pcs[i].UserID,
			&pcs[i].WorldID,
			&pcs[i].EntityID,
		); err != nil {
			return nil, nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return pcs[:i], state, nil
}

func (s Store) DeletePC(ctx context.Context, f entity.FilterPC) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.pc `)

	clause, args := filterPC(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
