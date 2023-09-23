package scylla

import (
	"context"
	"errors"
	"strings"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/gocql/gocql"
)

type filterPCPreferences entity.FilterPCPreferences

func (f filterPCPreferences) where() (string, []any) {
	var clause []string

	var args []any

	if f.PCID != nil {
		clause = append(clause, `pc_id = ?`)
		args = append(args, f.PCID)
	}

	if f.ID != nil {
		clause = append(clause, `id = ?`)
		args = append(args, f.ID)
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

func (f filterPCPreferences) index() string {
	var cols []string

	if f.PCID != nil {
		cols = append(cols, f.PCID.String())
	}

	if f.ID != nil {
		cols = append(cols, f.ID.String())
	}

	return strings.Join(cols, " - ")
}

type patchPCPreferences entity.PatchPCPreferences

func (p patchPCPreferences) set() (string, []any) {
	var clause []string

	var args []any

	if p.AbilityHotbars != nil {
		clause = append(clause, `ability_hotbars = ?`)
		args = append(args, p.AbilityHotbars)
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

func (s Store) InsertPCPreferences(ctx context.Context, pcp entity.PCPreferences) error {
	q := s.Session.Query(
		`INSERT INTO main.pc_preferences (id, pc_id, ability_hotbars) VALUES (?, ?, ?)`,
		pcp.ID,
		pcp.PCID,
		pcp.AbilityHotbars,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchPCPreferences(ctx context.Context, f entity.FilterPCPreferences) (entity.PCPreferences, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT id, pc_id, ability_hotbars FROM main.pc_preferences `)

	clause, args := filterPCPreferences(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var pcp entity.PCPreferences
	if err := q.Scan(&pcp.ID, &pcp.PCID, &pcp.AbilityHotbars); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return entity.PCPreferences{}, gerrors.ErrNotFound{Resource: "pc_preferences", Index: filterPCPreferences(f).index()}
		}

		return entity.PCPreferences{}, err
	}

	return pcp, nil
}

func (s Store) UpdatePCPreferences(
	ctx context.Context,
	f entity.FilterPCPreferences,
	p entity.PatchPCPreferences,
) error {
	b := strings.Builder{}
	b.WriteString(`UPDATE main.pc_preferences`)

	set, sArgs := patchPCPreferences(p).set()
	b.WriteString(set)

	clause, args := filterPCPreferences(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), append(sArgs, args...)...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) DeletePCPreferences(ctx context.Context, f entity.FilterPCPreferences) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.pc_preferences `)

	clause, args := filterPCPreferences(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
