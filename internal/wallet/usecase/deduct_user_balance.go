package usecase

import "github.com/paper-assessment/internal/models"

func (u *WalletUseCase) DeductUserBalance(request models.DeductUserBalanceRequest){
	res, _ := u.repo.GetUserWallet(request.UserId)

	remainingBalance := res.Balance - request.Amount

	if(remainingBalance >= 0) {
		u.repo.UpdateUserBalance(request.UserId, remainingBalance)
	}
}