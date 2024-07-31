package repository

import (
	"context"
	"errors"

	"github.com/paper-assessment/internal/wallet/database/schema"
	"github.com/paper-assessment/internal/wallet/domain/models"
	"gorm.io/gorm"
)

func (r *WalletRepository) Get(ctx context.Context, userId string) (*models.Wallet, error) {
	var wallet *schema.Wallet

	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &models.Wallet{
		UserId:  wallet.UserId,
		Balance: wallet.Balance,
	}, nil
}
