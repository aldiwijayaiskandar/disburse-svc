package handler

import (
	"encoding/json"
	"testing"

	"github.com/paper-assessment/internal/user/domain/models"
	"github.com/paper-assessment/internal/user/mocks"
	constants "github.com/paper-assessment/pkg/contants"
	"github.com/streadway/amqp"
)

func TestDeductBalanceHandler(t *testing.T) {
	mockPublisher := &mocks.MockPublisher{}
	mockUserUsecase := &mocks.MockUserUsecase{}
	replyTo := "reply_queue"
	request := models.GetUserRequest{
		Id: "1245",
	}

	handler := &RabbitMQHandler{
		publisher:   mockPublisher,
		userUsecase: mockUserUsecase,
	}

	t.Run("error request", func(t *testing.T) {
		mockPublisher.ExpectedPushAnyBody(replyTo)

		body, _ := json.Marshal(map[string]interface{}{
			"ids": request.Id,
		})

		handler.GetUserHandler(&amqp.Delivery{
			ReplyTo: replyTo,
			Body:    body,
		})

		mockPublisher.AssertNumberOfCalls(t, "Push", 1)
		mockUserUsecase.AssertNotCalled(t, "DeductBalance")

		mockPublisher.Reset()
	})

	t.Run("success", func(t *testing.T) {
		mockPublisher.ExpectedPushAnyBody(replyTo)
		mockUserUsecase.On("GetUser", request).Return(&models.GetUserResponse{
			Status: constants.Success,
			User: &models.User{
				Id:   request.Id,
				Name: "user",
			},
			ErrorCode: constants.NoError,
		})

		body, _ := json.Marshal(request)

		// call deduct balance
		handler.GetUserHandler(&amqp.Delivery{
			ReplyTo: replyTo,
			Body:    body,
		})

		// assert
		mockUserUsecase.AssertCalled(t, "GetUser", request)
		mockUserUsecase.AssertNumberOfCalls(t, "GetUser", 1)

		mockPublisher.Reset()
		mockUserUsecase.Reset()
	})
}
