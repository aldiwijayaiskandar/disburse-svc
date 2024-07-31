package repository

import "github.com/paper-assessment/internal/wallet/domain/models"

func (r *WalletRepository) Get(id string) (*models.Wallet, error) {
	return &models.Wallet{
		UserId:  "user-id",
		Balance: 5000,
	}, nil
}
