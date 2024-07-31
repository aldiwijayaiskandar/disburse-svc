package usecase

import (
	repository_interface "github.com/paper-assessment/internal/wallet/domain/repository"
)

type WalletUseCase struct {
	walletRepo repository_interface.WalletRepositoryInterface
}

type WalletUsecaseInterface interface {
}

func NewWalletUsecase(walletRepo repository_interface.WalletRepositoryInterface) WalletUsecaseInterface {
	return &WalletUseCase{
		walletRepo: walletRepo,
	}
}
