package repository

import (
	"context"

	"github.com/paper-assessment/internal/wallet/database/schema"
	"github.com/paper-assessment/internal/wallet/domain/models"
	"gorm.io/gorm"
)

func (r *WalletRepository) DeductBalance(ctx context.Context, request models.DeductBalanceRequest) error {
	result := r.db.WithContext(ctx).
		Model(&schema.Wallet{}).
		Where("user_id = ?", request.UserId).
		Update("balance", gorm.Expr("balance - ?", request.Amount))

	return result.Error
}
