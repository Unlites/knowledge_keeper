package models

type User struct {
	Id             uint
	Username       string
	PasswordHash   []byte
	RefreshToken   string
	TokenExpiresAt int64
}
