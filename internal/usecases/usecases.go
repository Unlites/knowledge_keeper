package usecases

import (
	"github.com/Unlites/knowledge_keeper/internal/infrastructure/repository"
)

type Auth interface {
}

type Record interface {
}

type Usecases struct {
	Auth
	Record
}

func NewUsecases(repo *repository.Repository) *Usecases {
	return &Usecases{
		Auth:   newAuthUsecase(repo.User),
		Record: newRecordUsecase(repo.Record),
	}
}
