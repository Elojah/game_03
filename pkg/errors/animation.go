package errors

import "fmt"

type ErrMissingEntity struct {
	AnimationID string
}

func (e ErrMissingEntity) Error() string {
	return fmt.Sprintf("missing required entity for animation %s", e.AnimationID)
}

type ErrMissingSheet struct {
	AnimationID string
}

func (e ErrMissingSheet) Error() string {
	return fmt.Sprintf("missing required sheet for animation %s", e.AnimationID)
}
