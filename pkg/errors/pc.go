package errors

import "fmt"

type ErrPCAlreadyConnected struct {
	ID string
}

func (e ErrPCAlreadyConnected) Error() string {
	return fmt.Sprintf("pc %s already connected", e.ID)
}

type ErrMissingDefaultAnimations struct {
	EntityID string
}

func (e ErrMissingDefaultAnimations) Error() string {
	return fmt.Sprintf("missing default animations for entity %s", e.EntityID)
}
