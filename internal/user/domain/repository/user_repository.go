package repository_interface

import (
	"context"

	"github.com/paper-assessment/internal/user/domain/models"
)

type UserRepositoryInterface interface {
	Get(ctx context.Context, id string) (*models.User, error)
}
