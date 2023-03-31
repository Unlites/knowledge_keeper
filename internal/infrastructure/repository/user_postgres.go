package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Unlites/knowledge_keeper/internal/errs"
	"github.com/Unlites/knowledge_keeper/internal/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	if err := ur.db.WithContext(ctx).Create(user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return &errs.ErrAlreadyExists{Object: "user with username " + user.Username}
		}
		return fmt.Errorf("failed to create user in db - %w", err)
	}

	return nil
}

func (ur *userRepository) GetUserByUsername(
	ctx context.Context,
	username string,
) (*models.User, error) {
	var user *models.User
	if err := ur.db.WithContext(ctx).
		First(&user, "username = ?", username).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errs.ErrNotFound{Object: "user with username " + username}
		}
		return nil, fmt.Errorf("failed to get user by username from db - %w", err)
	}

	return user, nil
}

func (ur *userRepository) GetUserByRefreshToken(
	ctx context.Context,
	refreshToken string,
) (*models.User, error) {
	var user *models.User
	if err := ur.db.WithContext(ctx).
		First(&user, "refresh_token = ?", refreshToken).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errs.ErrNotFound{Object: "user with such refresh_token"}
		}
		return nil, fmt.Errorf("failed to get user by refresh_token from db - %w", err)
	}

	return user, nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	if err := ur.db.WithContext(ctx).Save(user).Error; err != nil {
		return fmt.Errorf("failed to update user in db - %w", err)
	}

	return nil
}
