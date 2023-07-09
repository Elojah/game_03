package errors

import "fmt"

type ErrInvalidStat struct {
	Value string
}

func (e ErrInvalidStat) Error() string {
	return fmt.Sprintf("invalid stat %s", e.Value)
}

type ErrInvalidOperator struct {
	Value string
}

func (e ErrInvalidOperator) Error() string {
	return fmt.Sprintf("invalid operator %s", e.Value)
}

type ErrNotEnoughMP struct {
	Required  int64
	MP        int64
	AbilityID string
}

func (e ErrNotEnoughMP) Error() string {
	return fmt.Sprintf("not enough MP (%d/%d) to perform %s", e.MP, e.Required, e.AbilityID)
}

type ErrCooldownInProgress struct {
	At        int64
	LastCast  int64
	Cooldown  int64
	AbilityID string
}

func (e ErrCooldownInProgress) Error() string {
	return fmt.Sprintf("cooldown %d in progress (%d/%d) to perform %s", e.Cooldown, e.At, e.LastCast, e.AbilityID)
}

type ErrInvalidSelfTarget struct {
	AbilityID      string
	EffectID       string
	EffectTargetID string
	SourceID       string
	TargetID       string
}

func (e ErrInvalidSelfTarget) Error() string {
	return fmt.Sprintf(
		"invalid self target for ability %s effect %s target %s: received %s for source %s",
		e.AbilityID,
		e.EffectID,
		e.EffectTargetID,
		e.TargetID,
		e.SourceID,
	)
}

type ErrInvalidCastTargetFaction struct {
	AbilityID        string
	EffectID         string
	EffectTargetID   string
	EffectTargetType string
	SourceID         string
	SourceFactionID  string
	TargetID         string
	TargetFactionID  string
}

func (e ErrInvalidCastTargetFaction) Error() string {
	return fmt.Sprintf(
		"invalid target for ability %s effect %s target %s type %s: received %s faction %s for source %s faction %s",
		e.AbilityID,
		e.EffectID,
		e.EffectTargetID,
		e.EffectTargetType,
		e.TargetID,
		e.TargetFactionID,
		e.SourceID,
		e.SourceFactionID,
	)
}
