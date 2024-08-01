package services

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/paper-assessment/internal/disburse/mocks"
	"github.com/paper-assessment/internal/user/domain/models"
	constants "github.com/paper-assessment/pkg/contants"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {
	mockConsumer := &mocks.MockConsumer{}
	mockPublisher := &mocks.MockPublisher{}

	userService := &UserService{
		consumer:  mockConsumer,
		publisher: mockPublisher,
	}

	correlationId := uuid.New().String()
	request := &models.GetUserRequest{
		Id: uuid.New().String(),
	}
	response := &models.GetUserResponse{
		Status: constants.Success,
		User: &models.User{
			Id:   request.Id,
			Name: "user",
		},
	}
	resByte, _ := json.Marshal(response)

	t.Run("success", func(t *testing.T) {
		// mocking
		mockPublisher.ExpectedPushAnyBody("user.get.request", "user.get.request.reply", correlationId)
		mockConsumer.On("WaitReply", correlationId).Return(&amqp.Delivery{
			Body: resByte,
		}, nil)

		// call get user
		res := userService.GetUser(request, correlationId)

		// assertion
		assert.NotNil(t, res)
		assert.Equal(t, res.Status, constants.Success)
		assert.Equal(t, res.ErrorCode, constants.NoError)
		assert.Equal(t, res.User, response.User)

		mockPublisher.Reset()
		mockConsumer.Reset()
	})

	t.Run("consumer error", func(t *testing.T) {
		// mocks
		mockPublisher.ExpectedPushAnyBody("user.get.request", "user.get.request.reply", correlationId)
		mockConsumer.On("WaitReply", correlationId).Return(nil, errors.New("consumer_err"))

		// call get user
		res := userService.GetUser(request, correlationId)

		// assertion
		assert.NotNil(t, res)
		assert.Equal(t, res.Status, constants.Error)
		assert.Equal(t, res.ErrorCode, constants.InternalServerError)
		assert.Nil(t, res.User)

		mockPublisher.Reset()
		mockConsumer.Reset()
	})
}
