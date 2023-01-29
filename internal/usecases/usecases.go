package usecases

import (
	"context"

	"github.com/Unlites/knowledge_keeper/internal/dto"
	"github.com/Unlites/knowledge_keeper/internal/infrastructure/repository"
	"github.com/Unlites/knowledge_keeper/pkg/auth"
)

type Auth interface {
	SignUp(ctx context.Context, userDTO *dto.UserDTO) error
	SignIn(ctx context.Context, userDTO *dto.UserDTO) (*dto.TokensDTO, error)
	RefreshTokens(ctx context.Context, refreshToken *dto.RefreshTokenDTO) (*dto.TokensDTO, error)
}

type Record interface {
}

type Usecases struct {
	Auth
	Record
}

func NewUsecases(repo *repository.Repository, tokenManager auth.TokenManager, passwordHasher auth.PasswordHasher) *Usecases {
	return &Usecases{
		Auth:   newAuthUsecase(repo.User, tokenManager, passwordHasher),
		Record: newRecordUsecase(repo.Record),
	}
}
