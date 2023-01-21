package usecases

import (
	"github.com/Unlites/knowledge_keeper/internal/infrastructure/repository"
)

type recordUsecase struct {
	recordRepo repository.Record
}

func newRecordUsecase(recordRepo repository.Record) *recordUsecase {
	return &recordUsecase{
		recordRepo: recordRepo,
	}
}
