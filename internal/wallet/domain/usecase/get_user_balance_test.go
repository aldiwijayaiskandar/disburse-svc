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

func TestGetUserBalance(t *testing.T) {
	mockWalletRepo := &mocks.MockWalletRepository{}

	// expected variables
	userId := "user123"
	expectedBalance := 100.50

	usecase := &WalletUseCase{
		walletRepo: mockWalletRepo,
	}

	t.Run("success", func(t *testing.T) {
		// mocking wallet repo
		mockWalletRepo.On("Get", context.TODO(), userId).Return(&models.Wallet{
			UserId:  userId,
			Balance: expectedBalance,
		}, nil)

		// call usecase
		res := usecase.GetUserBalance(userId)

		// assert
		assert.NotNil(t, res.Balance)
		assert.Equal(t, constants.Success, res.Status)
		assert.Equal(t, expectedBalance, *res.Balance)
		assert.Equal(t, constants.NoError, res.ErrorCode)

		// reset mockRepo after each test
		mockWalletRepo.Reset()
	})

	t.Run("repo return not found", func(t *testing.T) {
		// mocking wallet repo
		mockWalletRepo.On("Get", context.TODO(), userId).Return(nil, nil)

		// call usecase
		res := usecase.GetUserBalance(userId)

		// assert
		assert.Equal(t, constants.Error, res.Status)
		assert.Equal(t, constants.NotFound, res.ErrorCode)

		// reset mockRepo after each test
		mockWalletRepo.Reset()
	})

	t.Run("repo return error", func(t *testing.T) {
		// mocking wallet repo
		mockWalletRepo.On("Get", context.TODO(), userId).Return(nil, errors.New("repo_error"))

		// call usecase
		res := usecase.GetUserBalance(userId)

		// assert
		assert.Equal(t, constants.Error, res.Status)
		assert.Equal(t, constants.InternalServerError, res.ErrorCode)

		// reset mockRepo after each test
		mockWalletRepo.Reset()
	})
}
