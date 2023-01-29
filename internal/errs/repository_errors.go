package errs

import "fmt"

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
