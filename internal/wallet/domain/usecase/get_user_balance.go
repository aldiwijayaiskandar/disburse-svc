package usecase

import (
	"context"

	"github.com/paper-assessment/internal/wallet/domain/models"
	constants "github.com/paper-assessment/pkg/contants"
)

func (u *WalletUseCase) GetUserBalance(userId string) *models.GetUserBalanceResponse {
	ctx := context.TODO()
	res, err := u.walletRepo.Get(ctx, userId)

	if err != nil {
		return &models.GetUserBalanceResponse{
			Status:    constants.Error,
			Balance:   nil,
			ErrorCode: constants.InternalServerError,
		}
	}

	if res == nil {
		return &models.GetUserBalanceResponse{
			Status:    constants.Error,
			Balance:   nil,
			ErrorCode: constants.NotFound,
		}
	}

	return &models.GetUserBalanceResponse{
		Status:  constants.Success,
		Balance: &res.Balance,
	}
}
