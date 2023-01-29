package errs

import (
	"errors"
)

var (
	ErrIncorrectPassword   = errors.New("incorrect password")
	ErrRefreshTokenExpired = errors.New("refresh token is expired")
)
