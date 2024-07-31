package repository

import (
	"context"

	"github.com/paper-assessment/internal/user/database/schema"
	"github.com/paper-assessment/internal/user/domain/models"
)

func (r UserRepository) Get(ctx context.Context, id string) (*models.User, error) {
	var user []schema.User

	if err := r.db.WithContext(ctx).Where("id = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}

	if len(user) == 0 {
		return nil, nil
	}

	return &models.User{
		Id:   user[0].Id,
		Name: user[0].Name,
	}, nil
}
