package usecase

import (
	"github.com/paper-assessment/internal/models"
)

func (u *UserUserCase) GetUser(req models.GetUserRequest) models.GetUserResponse {
	user, err := u.repo.GetById(req.Id)

	if err != nil {
		resErr := err.Error()
		return models.GetUserResponse{
			Error: &resErr,
		}
	}

	return models.GetUserResponse{
		User: &models.User{
			Id: user.Id,
			Name: user.Name,
		},
	}
}