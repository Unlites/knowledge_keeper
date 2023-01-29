package models

type User struct {
	Id             uint32
	Username       string
	PasswordHash   []byte
	RefreshToken   string
	TokenExpiresAt int64
}
