package scylla

import (
	"context"
	"errors"
	"strings"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/gocql/gocql"
)

type filterNPC entity.FilterNPC

func (f filterNPC) where() (string, []any) {
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

	b := strings.Builder{}
	b.WriteString(" WHERE ")

	if len(clause) == 0 {
		b.WriteString("false")
	} else {
		b.WriteString(strings.Join(clause, " AND "))
	}

	return b.String(), args
}

func (f filterNPC) index() string {
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

	return strings.Join(cols, " - ")
}

func (s Store) InsertNPC(ctx context.Context, npc entity.NPC) error {
	q := s.Session.Query(
		`INSERT INTO main.npc (id, world_id, entity_id) VALUES (?, ?, ?)`,
		npc.ID,
		npc.WorldID,
		npc.EntityID,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchNPC(ctx context.Context, f entity.FilterNPC) (entity.NPC, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT id, world_id, entity_id FROM main.npc `)

	clause, args := filterNPC(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var npc entity.NPC
	if err := q.Scan(&npc.ID, &npc.WorldID, &npc.EntityID); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return entity.NPC{}, gerrors.ErrNotFound{Resource: "npc", Index: filterNPC(f).index()}
		}

		return entity.NPC{}, err
	}

	return npc, nil
}

func (s Store) FetchManyNPC(ctx context.Context, f entity.FilterNPC) ([]entity.NPC, []byte, error) {
	if f.Size <= 0 {
		return nil, nil, nil
	}

	b := strings.Builder{}
	b.WriteString(`SELECT id, world_id, entity_id FROM main.npc `)

	clause, args := filterNPC(f).where()
	b.WriteString(clause)

	iter := s.Session.Query(b.String(), args...).
		WithContext(ctx).
		PageState(f.State).
		PageSize(f.Size).
		Iter()

	defer iter.Close()

	state := iter.PageState()

	scanner := iter.Scanner()

	pcs := make([]entity.NPC, f.Size)

	var i int

	for ; scanner.Next(); i++ {
		if err := scanner.Scan(
			&pcs[i].ID,
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

func (s Store) DeleteNPC(ctx context.Context, f entity.FilterNPC) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.npc `)

	clause, args := filterNPC(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
