package scylla

import (
	"context"
	"errors"
	"strings"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/gocql/gocql"
)

type filterAbility entity.FilterAbility

func (f filterAbility) where() (string, []any) {
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

	if f.EntityID != nil {
		clause = append(clause, `entity_id = ?`)
		args = append(args, f.EntityID)
	}

	if f.AbilityID != nil {
		clause = append(clause, `ability_id = ?`)
		args = append(args, f.AbilityID)
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

func (f filterAbility) index() string {
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

	if f.EntityID != nil {
		cols = append(cols, f.EntityID.String())
	}

	if f.AbilityID != nil {
		cols = append(cols, f.AbilityID.String())
	}

	return strings.Join(cols, " - ")
}

func (s Store) InsertAbility(ctx context.Context, ab entity.Ability) error {
	q := s.Session.Query(
		`INSERT INTO main.ab (entity_id, ability_id, last_cast) VALUES (?, ?, ?)`,
		ab.EntityID,
		ab.AbilityID,
		ab.LastCast,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchAbility(ctx context.Context, f entity.FilterAbility) (entity.Ability, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT entity_id, ability_id, last_cast FROM main.entity_ability `)

	clause, args := filterAbility(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var ab entity.Ability
	if err := q.Scan(&ab.EntityID, &ab.AbilityID, &ab.LastCast); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return entity.Ability{}, gerrors.ErrNotFound{Resource: "entity_ability", Index: filterAbility(f).index()}
		}

		return entity.Ability{}, err
	}

	return ab, nil
}

func (s Store) FetchManyAbility(ctx context.Context, f entity.FilterAbility) ([]entity.Ability, []byte, error) {
	if f.Size <= 0 {
		return nil, nil, nil
	}

	b := strings.Builder{}
	b.WriteString(`SELECT entity_id, ability_id, last_cast FROM main.entity_ability `)

	clause, args := filterAbility(f).where()
	b.WriteString(clause)

	iter := s.Session.Query(b.String(), args...).
		WithContext(ctx).
		PageState(f.State).
		PageSize(f.Size).
		Iter()

	defer iter.Close()

	state := iter.PageState()

	scanner := iter.Scanner()

	pcs := make([]entity.Ability, f.Size)

	var i int

	for ; scanner.Next(); i++ {
		if err := scanner.Scan(
			&pcs[i].EntityID,
			&pcs[i].AbilityID,
			&pcs[i].LastCast,
		); err != nil {
			return nil, nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return pcs[:i], state, nil
}

func (s Store) DeleteAbility(ctx context.Context, f entity.FilterAbility) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.entity_ability `)

	clause, args := filterAbility(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
