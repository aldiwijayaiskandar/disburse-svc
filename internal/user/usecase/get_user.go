package usecase

import (
	"github.com/google/uuid"
	"github.com/paper-assessment/internal/models"
)

func (u *UserUserCase) GetUser(req models.GetUserRequest) (*models.User) {
	return &models.User{
		Id: uuid.New().String(),
		Name: "user-test",
	}
}