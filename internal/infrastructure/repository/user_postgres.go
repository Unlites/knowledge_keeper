package repository

import (
	"context"
	"errors"
	"fmt"

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
	var userInDb *models.User
	result := ur.db.WithContext(ctx).First(&userInDb, "username = ?", user.Username)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &errs.ErrInternal{Message: fmt.Sprintf("failed to check user existence in db: %v", result.Error)}
	}

	if userInDb != nil {
		return &errs.ErrAlreadyExists{Object: "user with username " + user.Username}
	}

	result = ur.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return &errs.ErrInternal{Message: fmt.Sprintf("failed to create user in db: %v", result.Error)}
	}

	return nil
}
