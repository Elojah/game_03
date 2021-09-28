package errors

import "fmt"

type ErrPCAlreadyConnected struct {
	ID string
}

func (e ErrPCAlreadyConnected) Error() string {
	return fmt.Sprintf("pc %s already connected", e.ID)
}
