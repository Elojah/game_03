package errors

import "fmt"

type ErrNullRequest struct{}

func (e ErrNullRequest) Error() string {
	return fmt.Sprintf("null request")
}
