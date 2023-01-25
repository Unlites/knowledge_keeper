package usecases

import (
	"context"
	"fmt"

	"github.com/Unlites/knowledge_keeper/internal/dto"
	"github.com/Unlites/knowledge_keeper/internal/infrastructure/repository"
	"github.com/Unlites/knowledge_keeper/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	userRepo repository.User
}

func newAuthUsecase(userRepo repository.User) *authUsecase {
	return &authUsecase{
		userRepo: userRepo,
	}
}

func (au *authUsecase) SignUp(ctx context.Context, userDTO *dto.UserDTO) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), 10)
	if err != nil {
		return fmt.Errorf("failed to generate password hash: %v", err)
	}

	user := &models.User{
		Username:     userDTO.Username,
		PasswordHash: hash,
	}

	return au.userRepo.CreateUser(ctx, user)
}
