package usecase

import "github.com/paper-assessment/internal/models"

func (u *WalletUseCase) GetUserBalance(request models.GetUserBalanceRequest) (*models.GetUserBalanceResponse, error){
	res, err := u.repo.GetUserWallet(request.UserId)

	if err != nil {
		return nil, err
	}

	return &models.GetUserBalanceResponse{
		Balance: &res.Balance,
	}, nil
}