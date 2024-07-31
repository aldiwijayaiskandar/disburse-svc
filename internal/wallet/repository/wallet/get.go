package repository

import (
	"context"

	"github.com/paper-assessment/internal/wallet/database/schema"
	"github.com/paper-assessment/internal/wallet/domain/models"
)

func (r *WalletRepository) Get(ctx context.Context, userId string) (*models.Wallet, error) {
	var wallet []schema.Wallet

	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&wallet).Error; err != nil {
		return nil, err
	}

	if len(wallet) == 0 {
		return nil, nil
	}

	return &models.Wallet{
		UserId:  wallet[0].UserId,
		Balance: wallet[0].Balance,
	}, nil
}
