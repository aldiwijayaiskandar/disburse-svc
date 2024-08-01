package usecase

import (
	"context"

	"github.com/paper-assessment/internal/user/domain/models"
	constants "github.com/paper-assessment/pkg/contants"
)

func (u *UserUsecase) GetUser(request models.GetUserRequest) *models.GetUserResponse {
	ctx := context.TODO()

	res, err := u.userRepo.Get(ctx, request.Id)

	if err != nil {
		return &models.GetUserResponse{
			Status:    constants.Error,
			User:      nil,
			ErrorCode: constants.InternalServerError,
		}
	}

	if res == nil {
		return &models.GetUserResponse{
			Status:    constants.Error,
			User:      nil,
			ErrorCode: constants.NotFound,
		}
	}

	return &models.GetUserResponse{
		Status:    constants.Success,
		ErrorCode: constants.NoError,
		User:      res,
	}
}
