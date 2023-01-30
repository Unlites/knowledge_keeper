package usecases

import (
	"context"

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

	return ru.recordRepo.CreateRecord(ctx, userId, record)
}
