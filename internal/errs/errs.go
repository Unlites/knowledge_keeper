package errs

import "fmt"

type ErrValidation struct {
	Message string
}

func (e ErrValidation) Error() string {
	return fmt.Sprintf("validation error - %v", e.Message)
}

type ErrNotFound struct {
	Object string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("%s not found", e.Object)
}

type ErrAlreadyExists struct {
	Object string
}

func (e ErrAlreadyExists) Error() string {
	return fmt.Sprintf("%s already exists", e.Object)
}
