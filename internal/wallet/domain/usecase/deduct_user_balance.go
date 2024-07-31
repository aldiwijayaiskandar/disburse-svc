package usecase

import (
	"context"

	"github.com/paper-assessment/internal/wallet/domain/models"
	constants "github.com/paper-assessment/pkg/contants"
)

func (u *WalletUseCase) DeductBalance(request models.DeductBalanceRequest) *models.DeductBalanceResponse {
	ctx := context.TODO()
	res := u.getWallet(ctx, request.UserId)

	if res.ErrorCode != constants.NoError {
		return res
	}

	err := u.walletRepo.DeductBalance(ctx, request)

	if err != nil {
		return &models.DeductBalanceResponse{
			Status:    constants.Error,
			Balance:   nil,
			ErrorCode: constants.InternalServerError,
		}
	}

	return u.getWallet(ctx, request.UserId)
}

func (u *WalletUseCase) getWallet(ctx context.Context, userId string) *models.DeductBalanceResponse {
	res, err := u.walletRepo.Get(ctx, userId)

	if err != nil {
		return &models.DeductBalanceResponse{
			Status:    constants.Error,
			Balance:   nil,
			ErrorCode: constants.InternalServerError,
		}
	}

	if res == nil {
		return &models.DeductBalanceResponse{
			Status:    constants.Error,
			Balance:   nil,
			ErrorCode: constants.NotFound,
		}
	}

	return &models.DeductBalanceResponse{
		Status:    constants.Success,
		Balance:   &res.Balance,
		ErrorCode: constants.NoError,
	}
}
