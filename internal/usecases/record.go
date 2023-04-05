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

func (ru *recordUsecase) CreateRecord(
	ctx context.Context,
	userId uint,
	recordDTO *dto.RecordDTORequest,
) error {

	record := &models.Record{
		Topic:    recordDTO.Topic,
		Title:    recordDTO.Title,
		Subtopic: recordDTO.Subtopic,
		Content:  recordDTO.Content,
		UserId:   userId,
	}

	if err := ru.recordRepo.CreateRecord(ctx, record); err != nil {
		return fmt.Errorf("failed to create record - %w", err)
	}

	return nil
}

func (ru *recordUsecase) GetRecordById(
	ctx context.Context,
	userId, id uint,
) (*dto.RecordDTOResponse, error) {

	record, err := ru.recordRepo.GetRecordById(ctx, userId, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get record by id - %w", err)
	}

	return &dto.RecordDTOResponse{
		Id:       record.Id,
		Topic:    record.Topic,
		Subtopic: record.Subtopic,
		Title:    record.Title,
		Content:  record.Content,
	}, nil
}

func (ru *recordUsecase) GetAllRecords(
	ctx context.Context,
	userId uint,
	topic, subtopic, title string,
	offset, limit int,
) ([]*dto.RecordDTOResponse, error) {

	records, err := ru.recordRepo.GetAllRecords(
		ctx,
		userId,
		topic,
		subtopic,
		title,
		offset,
		limit,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get all records - %w", err)
	}

	recordDTOs := make([]*dto.RecordDTOResponse, 0)
	for _, record := range records {
		recordDTOs = append(recordDTOs, toDTO(record))
	}

	return recordDTOs, nil
}

func (ru *recordUsecase) GetAllTopics(
	ctx context.Context,
	userId uint,
) ([]string, error) {

	topics, err := ru.recordRepo.GetAllTopics(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get all topics - %w", err)
	}

	return topics, nil
}

func (ru *recordUsecase) GetAllSubtopicsByTopic(
	ctx context.Context,
	userId uint,
	topic string,
) ([]string, error) {

	topics, err := ru.recordRepo.GetAllSubtopicsByTopic(ctx, userId, topic)
	if err != nil {
		return nil, fmt.Errorf("failed to get subtopics - %w", err)
	}

	return topics, nil
}

func (ru *recordUsecase) UpdateRecord(
	ctx context.Context,
	userId, id uint,
	recordDTO *dto.RecordDTORequest,
) error {

	record, err := ru.recordRepo.GetRecordById(ctx, userId, id)
	if err != nil {
		return fmt.Errorf("failed to get updating record - %w", err)
	}

	record.Topic = recordDTO.Topic
	record.Subtopic = recordDTO.Subtopic
	record.Title = recordDTO.Title
	record.Content = recordDTO.Content

	if err := ru.recordRepo.UpdateRecord(ctx, record); err != nil {
		return fmt.Errorf("failed to update record - %w", err)
	}

	return nil
}

func (ru *recordUsecase) DeleteRecord(
	ctx context.Context,
	userId, id uint,
) error {

	record, err := ru.recordRepo.GetRecordById(ctx, userId, id)
	if err != nil {
		return fmt.Errorf("failed to get deleting record - %w", err)
	}

	if err := ru.recordRepo.DeleteRecord(ctx, record); err != nil {
		return fmt.Errorf("failed to delete record - %w", err)
	}

	return nil
}

func toDTO(record *models.Record) *dto.RecordDTOResponse {
	return &dto.RecordDTOResponse{
		Id:       record.Id,
		Topic:    record.Topic,
		Subtopic: record.Subtopic,
		Title:    record.Title,
		Content:  record.Content,
	}
}
