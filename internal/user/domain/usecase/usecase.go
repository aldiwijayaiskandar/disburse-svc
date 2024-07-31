package usecase

import (
	"github.com/paper-assessment/internal/user/domain/models"
	repository_interface "github.com/paper-assessment/internal/user/domain/repository"
)

type UserUsecase struct {
	userRepo repository_interface.UserRepositoryInterface
}

type UserUsecaseInterface interface {
	GetUser(request models.GetUserRequest) *models.GetUserResponse
}

func NewUserUsecase(userRepo repository_interface.UserRepositoryInterface) UserUsecaseInterface {
	return &UserUsecase{
		userRepo: userRepo,
	}
}
