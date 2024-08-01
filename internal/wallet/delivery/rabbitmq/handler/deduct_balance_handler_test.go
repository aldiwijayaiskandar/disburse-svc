package handler

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/paper-assessment/internal/wallet/domain/models"
	"github.com/paper-assessment/internal/wallet/mocks"
	constants "github.com/paper-assessment/pkg/contants"
	"github.com/streadway/amqp"
)

func TestDeductBalanceHandler(t *testing.T) {
	mockPublisher := &mocks.MockPublisher{}
	mockWalletUsecase := &mocks.WalletUsecase{}

	correlationId := uuid.New().String()
	replyTo := "reply_queue"
	request := models.DeductBalanceRequest{
		UserId: "1235",
		Amount: 5000.00,
	}

	handler := &RabbitMQHandler{
		publisher:     mockPublisher,
		walletUsecase: mockWalletUsecase,
	}

	t.Run("error request", func(t *testing.T) {
		mockPublisher.ExpectedReplyAnyBody(replyTo, correlationId)

		body, _ := json.Marshal(map[string]interface{}{
			"id": request.UserId,
		})

		handler.DeductBalanceHandler(&amqp.Delivery{
			ReplyTo:       replyTo,
			Body:          body,
			CorrelationId: correlationId,
		})

		mockWalletUsecase.AssertNotCalled(t, "DeductBalance")

		mockPublisher.Reset()
	})

	t.Run("success", func(t *testing.T) {
		mockPublisher.ExpectedReplyAnyBody(replyTo, correlationId)
		mockWalletUsecase.On("DeductBalance", request).Return(&models.DeductBalanceResponse{
			Status:    constants.Success,
			Balance:   &request.Amount,
			ErrorCode: constants.NoError,
		})

		body, _ := json.Marshal(request)

		// call deduct balance
		handler.DeductBalanceHandler(&amqp.Delivery{
			ReplyTo:       replyTo,
			Body:          body,
			CorrelationId: correlationId,
		})

		// assert
		mockWalletUsecase.AssertCalled(t, "DeductBalance", request)
		mockWalletUsecase.AssertNumberOfCalls(t, "DeductBalance", 1)

		mockPublisher.Reset()
	})
}
