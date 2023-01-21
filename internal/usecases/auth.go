package usecases

import (
	"github.com/Unlites/knowledge_keeper/internal/infrastructure/repository"
)

type authUsecase struct {
	userRepo repository.User
}

func newAuthUsecase(userRepo repository.User) *authUsecase {
	return &authUsecase{
		userRepo: userRepo,
	}
}
