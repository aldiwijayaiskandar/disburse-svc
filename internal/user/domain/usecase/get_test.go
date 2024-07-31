package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/paper-assessment/internal/user/domain/models"
	"github.com/paper-assessment/internal/user/mocks"
	constants "github.com/paper-assessment/pkg/contants"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	mockUserRepo := &mocks.MockUserRepository{}

	request := models.GetUserRequest{
		Id: "12345",
	}

	usecase := &UserUsecase{
		userRepo: mockUserRepo,
	}

	t.Run("success", func(t *testing.T) {
		response := &models.User{
			Id:   "12345",
			Name: "user",
		}

		// mocking wallet repo
		mockUserRepo.On("Get", context.TODO(), request.Id).Return(response, nil)

		// call usecase
		res := usecase.GetUser(request)

		// assert
		assert.NotNil(t, res)
		assert.Equal(t, constants.Success, res.Status)
		assert.Equal(t, response, res.User)
		assert.Equal(t, constants.NoError, res.ErrorCode)

		// reset mockRepo after each test
		mockUserRepo.Reset()
	})

	t.Run("repo return not found", func(t *testing.T) {
		// mocking wallet repo
		mockUserRepo.On("Get", context.TODO(), request.Id).Return(nil, nil)

		// call usecase
		res := usecase.GetUser(request)

		// assert
		assert.Equal(t, constants.Error, res.Status)
		assert.Equal(t, constants.NotFound, res.ErrorCode)

		// reset mockRepo after each test
		mockUserRepo.Reset()
	})

	t.Run("repo return error", func(t *testing.T) {
		// mocking wallet repo
		mockUserRepo.On("Get", context.TODO(), request.Id).Return(nil, errors.New("repo_error"))

		// call usecase
		res := usecase.GetUser(request)

		// assert
		assert.Equal(t, constants.Error, res.Status)
		assert.Equal(t, constants.InternalServerError, res.ErrorCode)

		// reset mockRepo after each test
		mockUserRepo.Reset()
	})
}
