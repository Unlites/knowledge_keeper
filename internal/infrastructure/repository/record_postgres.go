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

func (r *recordRepository) GetAllRecords(ctx context.Context, userId uint,
	topic string, offset, limit int) ([]*models.Record, error) {

	condition := map[string]interface{}{"user_id": userId}
	if topic != "" {
		condition["topic"] = topic
	}

	records := make([]*models.Record, 0)
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Where(condition).Find(&records).Error; err != nil {
		return nil, fmt.Errorf("failed to get records from db - %w", err)
	}
	return records, nil
}

func (r *recordRepository) GetAllTopics(ctx context.Context, userId uint) ([]string, error) {
	topics := make([]string, 0)
	if err := r.db.WithContext(ctx).Model(&models.Record{}).Distinct("topic").Find(&topics).Error; err != nil {
		return nil, fmt.Errorf("failed to get topics from db - %w", err)
	}

	return topics, nil
}

func (r *recordRepository) SearchRecordsByTitle(ctx context.Context,
	userId uint, title string, offset, limit int) ([]*models.Record, error) {

	records := make([]*models.Record, 0)

	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).
		Where("title iLIKE ?", "%"+title+"%").Find(&records).Error; err != nil {
		return nil, fmt.Errorf("failed to get records from db - %w", err)
	}

	return records, nil
}
