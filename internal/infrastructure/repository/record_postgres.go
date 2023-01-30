package repository

import (
	"context"
	"fmt"

	"github.com/Unlites/knowledge_keeper/internal/models"
	"gorm.io/gorm"
)

type recordRepository struct {
	db *gorm.DB
}

func newRecordRepository(db *gorm.DB) *recordRepository {
	return &recordRepository{
		db: db,
	}
}

func (r *recordRepository) CreateRecord(ctx context.Context, userId uint, record *models.Record) error {
	if err := r.db.WithContext(ctx).Create(record).Error; err != nil {
		return fmt.Errorf("failed to create user in db: %v", err)
	}

	return nil
}
