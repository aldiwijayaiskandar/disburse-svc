package usecase

import (
	"log"

	"github.com/google/uuid"
	"github.com/paper-assessment/internal/models"
)

func (u *UserUserCase) GetUser(req models.GetUserRequest) models.GetUserResponse {
	log.Println("testttt")

	return models.GetUserResponse{
		Data: &models.User{
			Id: uuid.New().String(),
			Name: "user-test",
		},
		Error: nil,
	}
}