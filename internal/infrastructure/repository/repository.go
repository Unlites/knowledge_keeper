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
	CreateRecord(ctx context.Context, record *models.Record) error
	GetRecordById(ctx context.Context, userId uint, id uint) (*models.Record, error)
	GetAllRecords(ctx context.Context, userId uint,
		topic, subtopic, title string, offset, limit int) ([]*models.Record, error)
	GetAllTopics(ctx context.Context, userId uint) ([]string, error)
	GetAllSubtopicsByTopic(ctx context.Context, userId uint, topic string) ([]string, error)
	UpdateRecord(ctx context.Context, record *models.Record) error
	DeleteRecord(ctx context.Context, record *models.Record) error
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
