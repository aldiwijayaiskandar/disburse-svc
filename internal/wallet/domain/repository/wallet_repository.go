package repository

import "github.com/paper-assessment/internal/wallet/domain/models"

type WalletRepository interface {
	Get(userId string) (*models.Wallet, error)
	Update(request models.UpdateBalanceRequest) (*models.Wallet, error)
}
