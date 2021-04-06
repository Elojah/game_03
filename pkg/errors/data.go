package errors

import "fmt"

type ErrNotFound struct {
	Resource string
	Index    string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("resource %s not found for %s", e.Resource, e.Index)
}
