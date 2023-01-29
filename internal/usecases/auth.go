package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/Unlites/knowledge_keeper/internal/dto"
	"github.com/Unlites/knowledge_keeper/internal/errs"
	"github.com/Unlites/knowledge_keeper/internal/infrastructure/repository"
	"github.com/Unlites/knowledge_keeper/internal/models"
	"github.com/Unlites/knowledge_keeper/pkg/auth"
)

type authUsecase struct {
	userRepo       repository.User
	tokenManager   auth.TokenManager
	passwordHasher auth.PasswordHasher
}

func newAuthUsecase(userRepo repository.User, tokenManager auth.TokenManager, passwordHasher auth.PasswordHasher) *authUsecase {
	return &authUsecase{
		userRepo:       userRepo,
		tokenManager:   tokenManager,
		passwordHasher: passwordHasher,
	}
}

func (au *authUsecase) SignUp(ctx context.Context, userDTO *dto.UserDTO) error {
	hash, err := au.passwordHasher.GenerateHash(userDTO.Password)
	if err != nil {
		return fmt.Errorf("failed to generate password hash: %v", err)
	}

	user := &models.User{
		Username:     userDTO.Username,
		PasswordHash: hash,
	}

	return au.userRepo.CreateUser(ctx, user)
}

func (au *authUsecase) SignIn(ctx context.Context, userDTO *dto.UserDTO) (*dto.TokensDTO, error) {
	user, err := au.userRepo.GetUserByUsername(ctx, userDTO.Username)
	if err != nil {
		return nil, err
	}

	if err := au.passwordHasher.Compare(userDTO.Password, user.PasswordHash); err != nil {
		return nil, errs.ErrIncorrectPassword
	}

	return au.updateUserTokens(ctx, user)
}

func (au *authUsecase) RefreshTokens(ctx context.Context, refreshToken *dto.RefreshTokenDTO) (*dto.TokensDTO, error) {
	user, err := au.userRepo.GetUserByRefreshToken(ctx, refreshToken.Token)
	if err != nil {
		return nil, err
	}

	if user.TokenExpiresAt < time.Now().Unix() {
		return nil, errs.ErrRefreshTokenExpired
	}

	return au.updateUserTokens(ctx, user)
}

func (au *authUsecase) updateUserTokens(ctx context.Context, user *models.User) (*dto.TokensDTO, error) {
	accessToken, err := au.tokenManager.NewAccessToken(string(rune(user.Id)))
	if err != nil {
		return nil, err
	}

	refreshToken, err := au.tokenManager.NewRefreshToken()
	if err != nil {
		return nil, err
	}

	user.RefreshToken = refreshToken.Token
	user.TokenExpiresAt = refreshToken.ExpiresAt

	if err := au.userRepo.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	return &dto.TokensDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken.Token,
	}, nil
}
