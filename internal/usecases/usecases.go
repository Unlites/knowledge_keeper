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
	SignOut(ctx context.Context, refreshToken *dto.RefreshTokenDTO) error
	RefreshTokens(ctx context.Context, refreshToken *dto.RefreshTokenDTO) (*dto.TokensDTO, error)
	ParseUserIdFromAccessToken(ctx context.Context, accessToken string) (string, error)
}

type Record interface {
	CreateRecord(ctx context.Context, userId uint, record *dto.RecordDTORequest) error
	GetRecordById(ctx context.Context, userId uint, id uint) (*dto.RecordDTOResponse, error)
	GetAllRecords(ctx context.Context, userId uint, topic string, offset, limit int) ([]*dto.RecordDTOResponse, error)
	GetAllTopics(ctx context.Context, userId uint) ([]string, error)
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
