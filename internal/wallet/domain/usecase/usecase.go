package usecase

import (
	repository_interface "github.com/paper-assessment/internal/wallet/domain/repository"
	repository "github.com/paper-assessment/internal/wallet/repository/wallet"
	"gorm.io/gorm"
)

type WalletUseCase struct{
	walletRepo repository_interface.WalletRepositoryInterface
}

type WalletUsecaseInterface interface {
}

func NewWalletUsecase(db *gorm.DB) WalletUsecaseInterface {
	walletRepo := repository.NewWalletRepository(db)

	return &WalletUseCase{
		walletRepo: walletRepo,
	}
}