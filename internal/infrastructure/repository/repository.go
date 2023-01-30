package repository

import (
	"context"

	"github.com/Unlites/knowledge_keeper/internal/models"
	"gorm.io/gorm"
)

type User interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetUserByRefreshToken(ctx context.Context, refreshToken string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
}

type Record interface {
	CreateRecord(ctx context.Context, userId uint, record *models.Record) error
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
