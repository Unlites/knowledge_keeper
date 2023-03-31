package usecases

import (
	"context"
	"fmt"
	"strconv"
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

func newAuthUsecase(
	userRepo repository.User,
	tokenManager auth.TokenManager,
	passwordHasher auth.PasswordHasher,
) *authUsecase {
	return &authUsecase{
		userRepo:       userRepo,
		tokenManager:   tokenManager,
		passwordHasher: passwordHasher,
	}
}

func (au *authUsecase) SignUp(ctx context.Context, userDTO *dto.UserDTO) error {
	hash, err := au.passwordHasher.GenerateHash(userDTO.Password)
	if err != nil {
		return fmt.Errorf("failed to generate password hash: %w", err)
	}

	user := &models.User{
		Username:     userDTO.Username,
		PasswordHash: hash,
	}

	if err := au.userRepo.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("failed to create user - %w", err)
	}

	return nil
}

func (au *authUsecase) SignIn(
	ctx context.Context,
	userDTO *dto.UserDTO,
) (*dto.TokensDTO, error) {
	user, err := au.userRepo.GetUserByUsername(ctx, userDTO.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by username - %w", err)
	}

	if err := au.passwordHasher.Compare(
		userDTO.Password,
		user.PasswordHash,
	); err != nil {
		return nil, errs.ErrIncorrectPassword
	}

	tokensDTO, err := au.updateUserTokens(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user tokens - %w", err)
	}

	return tokensDTO, nil
}

func (au *authUsecase) SignOut(
	ctx context.Context,
	refreshToken *dto.RefreshTokenDTO,
) error {
	user, err := au.userRepo.GetUserByRefreshToken(ctx, refreshToken.Token)
	if err != nil {
		return fmt.Errorf("failed to get user by refresh token - %w", err)
	}

	user.RefreshToken = ""
	user.TokenExpiresAt = 0

	if err := au.userRepo.UpdateUser(ctx, user); err != nil {
		return fmt.Errorf("failed to update user - %w", err)
	}

	return nil
}

func (au *authUsecase) RefreshTokens(
	ctx context.Context,
	refreshToken *dto.RefreshTokenDTO,
) (*dto.TokensDTO, error) {
	user, err := au.userRepo.GetUserByRefreshToken(ctx, refreshToken.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by refresh token - %w", err)
	}

	if user.TokenExpiresAt < time.Now().Unix() {
		return nil, errs.ErrRefreshTokenExpired
	}

	tokensDTO, err := au.updateUserTokens(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user tokens - %w", err)
	}

	return tokensDTO, nil
}

func (au *authUsecase) ParseUserIdFromAccessToken(
	ctx context.Context,
	accessToken string,
) (string, error) {
	userId, err := au.tokenManager.ParseUserIdFromAccessToken(accessToken)
	if err != nil {
		return "", fmt.Errorf("failed to parse user id from access token - %w", err)
	}

	return userId, nil
}

func (au *authUsecase) updateUserTokens(
	ctx context.Context,
	user *models.User,
) (*dto.TokensDTO, error) {
	accessToken, err := au.tokenManager.NewAccessToken(strconv.Itoa(int(user.Id)))
	if err != nil {
		return nil, fmt.Errorf("failed to create access token - %w", err)
	}

	refreshToken, err := au.tokenManager.NewRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("failed to create refresh token - %w", err)
	}

	user.RefreshToken = refreshToken.Token
	user.TokenExpiresAt = refreshToken.ExpiresAt

	if err := au.userRepo.UpdateUser(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to update user - %w", err)
	}

	return &dto.TokensDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken.Token,
	}, nil
}
