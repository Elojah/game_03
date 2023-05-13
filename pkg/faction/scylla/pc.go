package scylla

import (
	"context"
	"errors"
	"strings"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/faction"
	"github.com/gocql/gocql"
)

type filterPC faction.FilterPC

func (f filterPC) where() (string, []any) {
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

	if f.FactionID != nil {
		clause = append(clause, `faction_id = ?`)
		args = append(args, f.FactionID)
	}

	if len(f.FactionIDs) > 0 {
		clause = append(clause, `faction_id IN ?`)
		args = append(args, f.FactionIDs)
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

	if f.FactionID != nil {
		cols = append(cols, f.FactionID.String())
	}

	if f.FactionIDs != nil {
		ss := make([]string, 0, len(f.FactionIDs))
		for _, id := range f.FactionIDs {
			ss = append(ss, id.String())
		}

		cols = append(cols, strings.Join(ss, "|"))
	}

	return strings.Join(cols, " - ")
}

type patchPC faction.PatchPC

func (p patchPC) set() (string, []any) {
	var clause []string

	var args []any

	if p.FactionID != nil {
		clause = append(clause, `faction_id = ?`)
		args = append(args, p.FactionID)
	}

	if p.Permission != nil {
		clause = append(clause, `permission = ?`)
		args = append(args, *p.Permission)
	}

	b := strings.Builder{}
	b.WriteString(" SET ")

	if len(clause) == 0 {
		b.WriteString("0") // raise error update
	} else {
		b.WriteString(strings.Join(clause, ", "))
	}

	return b.String(), args
}

func (s Store) InsertPC(ctx context.Context, pc faction.PC) error {
	q := s.Session.Query(
		`INSERT INTO main.faction_pc (id, faction_id, permission) VALUES (?, ?, ?)`,
		pc.ID,
		pc.FactionID,
		pc.Permission,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) UpdatePC(ctx context.Context, f faction.FilterPC, p faction.PatchPC) error {
	b := strings.Builder{}
	b.WriteString(`UPDATE main.faction_pc`)

	set, sArgs := patchPC(p).set()
	b.WriteString(set)

	clause, args := filterPC(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), append(sArgs, args...)...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchPC(ctx context.Context, f faction.FilterPC) (faction.PC, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT id, faction_id, permission FROM main.faction_pc `)

	clause, args := filterPC(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var pc faction.PC
	if err := q.Scan(&pc.ID, &pc.FactionID, &pc.Permission); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return faction.PC{}, gerrors.ErrNotFound{Resource: "faction_pc", Index: filterPC(f).index()}
		}

		return faction.PC{}, err
	}

	return pc, nil
}

func (s Store) FetchManyPC(ctx context.Context, f faction.FilterPC) ([]faction.PC, []byte, error) {
	if f.Size <= 0 {
		return nil, nil, nil
	}

	b := strings.Builder{}
	b.WriteString(`SELECT id, faction_id, permission FROM main.faction_pc `)

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

	pcs := make([]faction.PC, f.Size)

	var i int

	for ; scanner.Next(); i++ {
		if err := scanner.Scan(
			&pcs[i].ID,
			&pcs[i].FactionID,
			&pcs[i].Permission,
		); err != nil {
			return nil, nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return pcs[:i], state, nil
}

func (s Store) DeletePC(ctx context.Context, f faction.FilterPC) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.faction_pc `)

	clause, args := filterPC(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
