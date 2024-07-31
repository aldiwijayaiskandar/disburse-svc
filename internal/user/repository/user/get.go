package repository

import (
	"context"
	"errors"

	"github.com/paper-assessment/internal/user/database/schema"
	"github.com/paper-assessment/internal/user/domain/models"
	"gorm.io/gorm"
)

func (r UserRepository) Get(ctx context.Context, id string) (*models.User, error) {
	var user schema.User

	if err := r.db.WithContext(ctx).Where("id = ?", id).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &models.User{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}
