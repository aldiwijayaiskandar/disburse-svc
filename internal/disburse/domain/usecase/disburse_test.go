package usecase

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/paper-assessment/internal/disburse/domain/models"
	"github.com/paper-assessment/internal/disburse/mocks"
	usermodel "github.com/paper-assessment/internal/user/domain/models"
	walletmodel "github.com/paper-assessment/internal/wallet/domain/models"
	constants "github.com/paper-assessment/pkg/contants"
	"github.com/stretchr/testify/assert"
)

func Test_Disburse(t *testing.T) {
	mockUserService := &mocks.MockUserService{}
	mockWalletService := &mocks.MockWalletService{}

	usecase := &DisburseUsecase{
		userService:   mockUserService,
		walletService: mockWalletService,
	}

	correlationId := uuid.New().String()

	userId := "user-id"
	amount := 5000.00

	disburseRequest := models.DisburseRequest{
		UserId: userId,
		Amount: amount,
	}

	getUserRequest := usermodel.GetUserRequest{
		Id: userId,
	}
	getUserSuccessResponse := usermodel.GetUserResponse{
		Status:    constants.Success,
		ErrorCode: constants.NoError,
		User: &usermodel.User{
			Id:   userId,
			Name: "user",
		},
	}

	getUserBalanceRequest := walletmodel.GetUserBalanceRequest{
		UserId: userId,
	}
	getUserBalanceSuccessResponse := walletmodel.GetUserBalanceResponse{
		Status:    constants.Success,
		ErrorCode: constants.NoError,
		Balance:   &amount,
	}

	deductUserBalanceRequest := walletmodel.DeductBalanceRequest{
		UserId: userId,
		Amount: amount,
	}
	deductUserBalanceSuccessResponse := walletmodel.DeductBalanceResponse{
		Status:    constants.Success,
		ErrorCode: constants.NoError,
		Balance:   &amount,
	}

	testErrorMessage := errors.New("test-error").Error()

	t.Run("success", func(t *testing.T) {
		// mock get user to return error
		mockUserService.
			On("GetUser", &getUserRequest, correlationId).
			Return(getUserSuccessResponse)
		// mock get balance success
		mockWalletService.
			On("GetUserBalance", &getUserBalanceRequest, correlationId).
			Return(getUserBalanceSuccessResponse)
		// mock deduct user balance success
		mockWalletService.
			On("DeductUserBalance", &deductUserBalanceRequest, correlationId).
			Return(deductUserBalanceSuccessResponse)

		// calling disburse
		res := usecase.Disburse(disburseRequest, correlationId)

		// assertion
		assert.Equal(t, deductUserBalanceSuccessResponse.Status, res.Status)
		assert.Equal(t, deductUserBalanceSuccessResponse.ErrorCode, res.ErrorCode)
		assert.Equal(t, deductUserBalanceSuccessResponse.Message, res.Message)
		assert.Equal(t, deductUserBalanceSuccessResponse.Balance, res.Balance)

		mockUserService.Reset()
		mockWalletService.Reset()
	})

	t.Run("get user error", func(t *testing.T) {
		expectedResponse := usermodel.GetUserResponse{
			Status:    constants.Error,
			ErrorCode: constants.InvalidRequest,
			Message:   &testErrorMessage,
		}

		// mock get user to return error
		mockUserService.
			On("GetUser", &getUserRequest, correlationId).
			Return(
				expectedResponse,
			)

		// calling disburse
		res := usecase.Disburse(disburseRequest, correlationId)

		// assertion
		assert.Equal(t, expectedResponse.Status, res.Status)
		assert.Equal(t, expectedResponse.ErrorCode, res.ErrorCode)
		assert.Equal(t, expectedResponse.Message, res.Message)

		// clean up
		mockUserService.Reset()
	})

	t.Run("get wallet error", func(t *testing.T) {
		expectedResponse := walletmodel.GetUserBalanceResponse{
			Status:    constants.Error,
			ErrorCode: constants.InvalidRequest,
			Message:   &testErrorMessage,
		}

		// mock get user to return error
		mockUserService.
			On("GetUser", &getUserRequest, correlationId).
			Return(getUserSuccessResponse)
		// mock get balance success
		mockWalletService.
			On("GetUserBalance", &getUserBalanceRequest, correlationId).
			Return(expectedResponse)

		// calling disburse
		res := usecase.Disburse(disburseRequest, correlationId)

		// assertion
		assert.Equal(t, expectedResponse.Status, res.Status)
		assert.Equal(t, expectedResponse.ErrorCode, res.ErrorCode)
		assert.Equal(t, expectedResponse.Message, res.Message)

		// clean up
		mockUserService.Reset()
		mockWalletService.Reset()
	})

	t.Run("balance insufficient", func(t *testing.T) {
		req := models.DisburseRequest{
			UserId: disburseRequest.UserId,
			Amount: amount * 2,
		}

		// mock get user to return error
		mockUserService.
			On("GetUser", &getUserRequest, correlationId).
			Return(getUserSuccessResponse)
		// mock get balance success
		mockWalletService.
			On("GetUserBalance", &getUserBalanceRequest, correlationId).
			Return(getUserBalanceSuccessResponse)
		// mock deduct user balance success
		mockWalletService.
			On("DeductUserBalance", &deductUserBalanceRequest, correlationId).
			Return(deductUserBalanceSuccessResponse)

		// calling disburse
		res := usecase.Disburse(req, correlationId)

		// assertion
		assert.Equal(t, constants.Error, res.Status)
		assert.Equal(t, constants.InvalidRequest, res.ErrorCode)
		assert.Equal(t, "balance insufficient", *res.Message)

		// clean up
		mockUserService.Reset()
		mockWalletService.Reset()
	})
}
