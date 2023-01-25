package models

type User struct {
	Id           uint32
	Username     string
	PasswordHash []byte
	Records      []Record
}
