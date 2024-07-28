package repository

import (
	"github.com/paper-assessment/internal/wallet/models"
	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (u *WalletRepository) GetUserBalance(userId string) (float64, error) {
    var wallet models.Wallet

    if result := u.db.Where(&models.Wallet{UserId: userId}).First(&wallet); result.Error != nil {
        return 0, result.Error
    }

    return wallet.Balance, nil
}

func (u *WalletRepository) DeductUserBalance(userId string, amount float64) {
    
}