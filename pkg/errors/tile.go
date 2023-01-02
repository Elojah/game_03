package errors

import "fmt"

type ErrMissingWangSet struct {
	ID string
}

func (err ErrMissingWangSet) Error() string {
	return fmt.Sprintf("missing or invalid wangset in tileset %s", err.ID)
}

type ErrMissingCollisionLayer struct {
	ID string
}

func (err ErrMissingCollisionLayer) Error() string {
	return fmt.Sprintf("missing or invalid collision layer in tileset %s", err.ID)
}
