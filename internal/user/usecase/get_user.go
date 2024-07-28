package usecase

import (
	"github.com/paper-assessment/internal/models"
)

func (u *UserUserCase) GetUser(req models.GetUserRequest) models.GetUserResponse {
	user, err := u.repo.GetById(req.Id)

	if err != nil {
		panic(err)
	}

	return models.GetUserResponse{
		Data: &models.User{
			Id: user.Id,
			Name: user.Name,
		},
		Error: nil,
	}
}