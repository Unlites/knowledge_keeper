package repository

import "gorm.io/gorm"

type User interface {
}

type Record interface {
}

type Repository struct {
	User
	Record
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:   newUserRepository(db),
		Record: newRecordRepository(db),
	}
}
