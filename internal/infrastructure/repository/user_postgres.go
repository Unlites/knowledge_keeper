package repository

import "gorm.io/gorm"

type userRepository struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}
