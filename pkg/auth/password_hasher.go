package auth

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher interface {
	Compare(password string, hash []byte) error
	GenerateHash(password string) ([]byte, error)
}

type PasswordHasherSettings struct {
	Cost int
}

type hasher struct {
	settings *PasswordHasherSettings
}

var _ PasswordHasher = (*hasher)(nil)

func NewPasswordHasher(settings *PasswordHasherSettings) *hasher {
	return &hasher{settings: settings}
}

func (h *hasher) GenerateHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), h.settings.Cost)
}

func (h *hasher) Compare(password string, hash []byte) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(password))
}
