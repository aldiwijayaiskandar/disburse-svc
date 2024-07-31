package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/paper-assessment/internal/wallet/domain/models"
	"github.com/paper-assessment/internal/wallet/mocks"
	constants "github.com/paper-assessment/pkg/contants"
	"github.com/stretchr/testify/assert"
)

func TestDeductUserBalance(t *testing.T) {
	mockWalletRepo := &mocks.MockWalletRepository{}

	// expected variables
	request := models.DeductBalanceRequest{
		UserId: "12345",
		Amount: 5000.00,
	}

	usecase := &WalletUseCase{
		walletRepo: mockWalletRepo,
	}

	t.Run("success", func(t *testing.T) {
		// mocking wallet repo
		mockWalletRepo.On("Get", context.TODO(), request.UserId).Return(&models.Wallet{
			UserId:  request.UserId,
			Balance: request.Amount,
		}, nil)
		mockWalletRepo.On("DeductBalance", context.TODO(), request).Return(nil)

		// call usecase
		res := usecase.DeductBalance(request)

		// assert
		assert.NotNil(t, res)
		assert.Equal(t, res.Status, constants.Success)
		assert.Equal(t, res.ErrorCode, constants.NoError)
		assert.Nil(t, res.Message)
		assert.Equal(t, res.Balance, &request.Amount)

		// reset mockRepo after each test
		mockWalletRepo.Reset()
	})

	t.Run("repo return not found", func(t *testing.T) {
		// mocking wallet repo
		mockWalletRepo.On("Get", context.TODO(), request.UserId).Return(nil, nil)

		// call usecase
		res := usecase.DeductBalance(request)

		// assert
		assert.Equal(t, constants.Error, res.Status)
		assert.Equal(t, constants.NotFound, res.ErrorCode)

		// reset mockRepo after each test
		mockWalletRepo.Reset()
	})

	t.Run("repo return error", func(t *testing.T) {
		// mocking wallet repo
		mockWalletRepo.On("Get", context.TODO(), request.UserId).Return(nil, errors.New("repo_error"))

		// call usecase
		res := usecase.DeductBalance(request)

		// assert
		assert.Equal(t, constants.Error, res.Status)
		assert.Equal(t, constants.InternalServerError, res.ErrorCode)

		// reset mockRepo after each test
		mockWalletRepo.Reset()
	})
}
