package usecase

import (
	"github.com/paper-assessment/internal/wallet/domain/models"
	repository_interface "github.com/paper-assessment/internal/wallet/domain/repository"
)

type WalletUseCase struct {
	walletRepo repository_interface.WalletRepositoryInterface
}

type WalletUsecaseInterface interface {
	GetUserBalance(id string) *models.GetUserBalanceResponse
	DeductBalance(request models.DeductBalanceRequest) *models.DeductBalanceResponse
}

func NewWalletUsecase(walletRepo repository_interface.WalletRepositoryInterface) WalletUsecaseInterface {
	return &WalletUseCase{
		walletRepo: walletRepo,
	}
}
