package repository

import (
	"context"

	"github.com/Unlites/knowledge_keeper/internal/models"
	"gorm.io/gorm"
)

type User interface {
	CreateUser(ctx context.Context, user *models.User) error
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
