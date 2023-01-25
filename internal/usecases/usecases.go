package usecases

import (
	"context"

	"github.com/Unlites/knowledge_keeper/internal/dto"
	"github.com/Unlites/knowledge_keeper/internal/infrastructure/repository"
)

type Auth interface {
	SignUp(ctx context.Context, userDTO *dto.UserDTO) error
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
