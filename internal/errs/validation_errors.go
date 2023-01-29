package errs

import "fmt"

type ErrValidation struct {
	Message string
}

func (e ErrValidation) Error() string {
	return fmt.Sprintf("validation error - %v", e.Message)
}
