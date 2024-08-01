package services

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/paper-assessment/internal/disburse/mocks"
	"github.com/paper-assessment/internal/wallet/domain/models"
	constants "github.com/paper-assessment/pkg/contants"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func TestWalletService_GetUserBalance(t *testing.T) {
	mockConsumer := &mocks.MockConsumer{}
	mockPublisher := &mocks.MockPublisher{}

	walletService := &WalletService{
		consumer:  mockConsumer,
		publisher: mockPublisher,
	}

	correlationId := uuid.New().String()

	t.Run("get user balance", func(t *testing.T) {
		balance := 5000.00
		request := &models.GetUserBalanceRequest{
			UserId: uuid.New().String(),
		}
		response := &models.GetUserBalanceResponse{
			Status:  constants.Success,
			Balance: &balance,
		}
		resByte, _ := json.Marshal(response)

		t.Run("success", func(t *testing.T) {
			// mocking
			mockPublisher.ExpectedPushAnyBody("wallet.balance.get.request", correlationId)
			mockConsumer.On("WaitReply", correlationId).Return(&amqp.Delivery{
				Body: resByte,
			}, nil)

			// call get user
			res := walletService.GetUserBalance(request, correlationId)

			// assertion
			assert.NotNil(t, res)
			assert.Equal(t, res.Status, constants.Success)
			assert.Equal(t, res.ErrorCode, constants.NoError)
			assert.Equal(t, res.Balance, response.Balance)

			mockPublisher.Reset()
			mockConsumer.Reset()
		})

		t.Run("consumer error", func(t *testing.T) {
			// mocks
			mockPublisher.ExpectedPushAnyBody("wallet.balance.get.request", correlationId)
			mockConsumer.On("WaitReply", correlationId).Return(nil, errors.New("consumer_err"))

			// call get user
			res := walletService.GetUserBalance(request, correlationId)

			// assertion
			assert.NotNil(t, res)
			assert.Equal(t, res.Status, constants.Error)
			assert.Equal(t, res.ErrorCode, constants.InternalServerError)
			assert.Nil(t, res.Balance)

			mockPublisher.Reset()
			mockConsumer.Reset()
		})
	})

	t.Run("deduct balance", func(t *testing.T) {
		balance := 5000.00
		request := &models.DeductBalanceRequest{
			UserId: uuid.New().String(),
			Amount: 500.00,
		}
		response := &models.DeductBalanceResponse{
			Status:  constants.Success,
			Balance: &balance,
		}
		resByte, _ := json.Marshal(response)

		t.Run("success", func(t *testing.T) {
			mockPublisher.ExpectedPushAnyBody("wallet.balance.deduct.request", correlationId)
			mockConsumer.On("WaitReply", correlationId).Return(&amqp.Delivery{
				Body: resByte,
			}, nil)

			// call get user
			res := walletService.DeductUserBalance(request, correlationId)

			// assertion
			assert.NotNil(t, res)
			assert.Equal(t, res.Status, constants.Success)
			assert.Equal(t, res.ErrorCode, constants.NoError)
			assert.Equal(t, res.Balance, response.Balance)

			mockPublisher.Reset()
			mockConsumer.Reset()
		})

		t.Run("consumer error", func(t *testing.T) {
			mockPublisher.ExpectedPushAnyBody("wallet.balance.deduct.request", correlationId)
			mockConsumer.On("WaitReply", correlationId).Return(nil, errors.New("consumer_err"))

			// call get user
			res := walletService.DeductUserBalance(request, correlationId)

			// assertion
			assert.NotNil(t, res)
			assert.Equal(t, res.Status, constants.Error)
			assert.Equal(t, res.ErrorCode, constants.InternalServerError)
			assert.Nil(t, res.Balance)

			mockPublisher.Reset()
			mockConsumer.Reset()
		})
	})
}
