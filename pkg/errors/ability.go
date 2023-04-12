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
