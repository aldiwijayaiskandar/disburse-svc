package repository

import "github.com/paper-assessment/internal/wallet/domain/repository"

type WalletRepository struct {
}

func NewWalletRepository() repository.WalletRepository {
	return &WalletRepository{}
}
