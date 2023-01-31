package usecases

import (
	"context"
	"fmt"

	"github.com/Unlites/knowledge_keeper/internal/dto"
	"github.com/Unlites/knowledge_keeper/internal/infrastructure/repository"
	"github.com/Unlites/knowledge_keeper/internal/models"
)

type recordUsecase struct {
	recordRepo repository.Record
}

func newRecordUsecase(recordRepo repository.Record) *recordUsecase {
	return &recordUsecase{
		recordRepo: recordRepo,
	}
}

func (ru *recordUsecase) CreateRecord(ctx context.Context, userId uint, recordDTO *dto.RecordDTORequest) error {
	record := &models.Record{
		Topic:   recordDTO.Topic,
		Title:   recordDTO.Title,
		Content: recordDTO.Content,
		UserId:  userId,
	}

	if err := ru.recordRepo.CreateRecord(ctx, userId, record); err != nil {
		return fmt.Errorf("failed to create record - %w", err)
	}

	return nil
}

func (ru *recordUsecase) GetRecordById(ctx context.Context, userId uint, id uint) (*dto.RecordDTOResponse, error) {
	record, err := ru.recordRepo.GetRecordById(ctx, userId, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get record by id - %w", err)
	}

	return &dto.RecordDTOResponse{
		Id:      record.Id,
		Topic:   record.Topic,
		Title:   record.Title,
		Content: record.Content,
	}, nil
}

func (ru *recordUsecase) GetAllRecords(ctx context.Context, userId uint, offset, limit int) ([]*dto.RecordDTOResponse, error) {
	return nil, nil
}

func (ru *recordUsecase) GetAllRecordsByTopic(ctx context.Context, userId uint,
	topic string, offset, limit int) ([]*dto.RecordDTOResponse, error) {
	return nil, nil
}
