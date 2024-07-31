package repository_interface

import (
	"context"

	"github.com/paper-assessment/internal/wallet/domain/models"
)

type WalletRepositoryInterface interface {
	Get(ctx context.Context, userId string) (*models.Wallet, error)
	DeductBalance(ctx context.Context, request models.DeductBalanceRequest) error
}
