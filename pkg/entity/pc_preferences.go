package entity

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type FilterPCPreferences struct {
	ID   ulid.ID
	PCID ulid.ID
}

type PatchPCPreferences struct {
	AbilityHotbars map[string]ulid.ID
}

type StorePCPreferences interface {
	InsertPCPreferences(context.Context, PCPreferences) error
	FetchPCPreferences(context.Context, FilterPCPreferences) (PCPreferences, error)
	UpdatePCPreferences(context.Context, FilterPCPreferences, PatchPCPreferences) error
	DeletePCPreferences(context.Context, FilterPCPreferences) error
}
