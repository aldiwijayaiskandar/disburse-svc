package usecase

import (
	"github.com/paper-assessment/internal/models"
)

func (u *UserUserCase) GetUser(req models.GetUserRequest) models.GetUserResponse {
	user := u.repo.GetById(req.Id)

	return models.GetUserResponse{
		Data: &models.User{
			Id: user.Id,
			Name: user.Name,
		},
		Error: nil,
	}
}