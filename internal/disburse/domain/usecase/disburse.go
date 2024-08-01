package usecase

import (
	"errors"

	"github.com/paper-assessment/internal/disburse/domain/models"
	usermodel "github.com/paper-assessment/internal/user/domain/models"
	walletmodel "github.com/paper-assessment/internal/wallet/domain/models"
	constants "github.com/paper-assessment/pkg/contants"
)

func (u *DisburseUsecase) Disburse(request models.DisburseRequest, correlationId string) *models.DisburseResponse {
	// check user exist
	userRes := u.userService.GetUser(&usermodel.GetUserRequest{
		Id: request.UserId,
	}, correlationId)

	if userRes.Status == constants.Error {
		// throw user error code
		return &models.DisburseResponse{
			Status:    userRes.Status,
			ErrorCode: userRes.ErrorCode,
			Message:   userRes.Message,
		}
	}

	// get user balance
	balanceRes := u.walletService.GetUserBalance(&walletmodel.GetUserBalanceRequest{
		UserId: request.UserId,
	}, correlationId)

	if balanceRes.Status == constants.Error {
		// throw wallet error code
		return &models.DisburseResponse{
			Status:    balanceRes.Status,
			ErrorCode: balanceRes.ErrorCode,
			Message:   balanceRes.Message,
		}
	}

	if balanceRes.Balance != nil && *balanceRes.Balance < request.Amount {
		// throw insufficient amount
		errorMessage := errors.New("balance insufficient").Error()

		return &models.DisburseResponse{
			ErrorCode: constants.InvalidRequest,
			Status:    constants.Error,
			Message:   &errorMessage,
		}
	}

	deductBalanceRes := u.walletService.DeductUserBalance(&walletmodel.DeductBalanceRequest{
		UserId: request.UserId,
		Amount: request.Amount,
	}, correlationId)

	return &models.DisburseResponse{
		Status:    deductBalanceRes.Status,
		Balance:   deductBalanceRes.Balance,
		ErrorCode: deductBalanceRes.ErrorCode,
		Message:   deductBalanceRes.Message,
	}
}
