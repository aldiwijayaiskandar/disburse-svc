package handler

import (
	"encoding/json"
	"testing"

	"github.com/paper-assessment/internal/wallet/domain/models"
	"github.com/paper-assessment/internal/wallet/mocks"
	constants "github.com/paper-assessment/pkg/contants"
	"github.com/streadway/amqp"
)

func TestGetUserBalanceHandler(t *testing.T) {
	mockPublisher := &mocks.MockPublisher{}
	mockWalletUsecase := &mocks.WalletUsecase{}
	replyTo := "reply_queue"
	expectedUserId := "12345"
	expectedBalance := 5000.00

	handler := &RabbitMQHandler{
		publisher:     mockPublisher,
		walletUsecase: mockWalletUsecase,
	}

	t.Run("error request", func(t *testing.T) {
		mockPublisher.ExpectedPushAnyBody(replyTo)

		body, _ := json.Marshal(map[string]interface{}{
			"id": expectedUserId,
		})

		handler.GetUserBalanceHandler(&amqp.Delivery{
			ReplyTo: replyTo,
			Body:    body,
		})

		mockWalletUsecase.AssertNotCalled(t, "GetUserBalance")

		mockPublisher.Reset()
	})

	t.Run("success", func(t *testing.T) {
		mockPublisher.ExpectedPushAnyBody(replyTo)
		mockWalletUsecase.On("GetUserBalance", expectedUserId).Return(&models.GetUserBalanceResponse{
			Status:    constants.Success,
			Balance:   &expectedBalance,
			ErrorCode: constants.NoError,
		})

		body, _ := json.Marshal(map[string]interface{}{
			"userId": expectedUserId,
		})

		handler.GetUserBalanceHandler(&amqp.Delivery{
			ReplyTo: replyTo,
			Body:    body,
		})

		mockWalletUsecase.AssertCalled(t, "GetUserBalance", expectedUserId)

		mockPublisher.Reset()
	})
}
