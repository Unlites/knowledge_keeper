package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Unlites/knowledge_keeper/internal/errs"
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
		return fmt.Errorf("failed to create user in db - %w", err)
	}

	return nil
}

func (r *recordRepository) GetRecordById(ctx context.Context, userId uint, id uint) (*models.Record, error) {
	var record *models.Record
	if err := r.db.WithContext(ctx).First(&record, "user_id = ? AND id = ?", userId, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errs.ErrNotFound{Object: "record"}
		}
		return nil, fmt.Errorf("failed to get record by id in db - %w", err)
	}

	return record, nil
}
func (r *recordRepository) GetAllRecords(ctx context.Context, userId uint, offset, limit int) ([]*models.Record, error) {
	return nil, nil
}
func (r *recordRepository) GetAllRecordsByTopic(ctx context.Context, userId uint,
	topic string, offset, limit int) ([]*models.Record, error) {
	return nil, nil
}
