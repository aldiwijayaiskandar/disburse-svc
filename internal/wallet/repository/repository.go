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

func (u *WalletRepository) GetUserWallet(userId string) (*models.Wallet, error) {
    var wallet *models.Wallet

    if result := u.db.Where(&models.Wallet{UserId: userId}).First(&wallet); result.Error != nil {
        return nil, result.Error
    }

    return wallet, nil
}

func (u *WalletRepository) DeductUserBalance(userId string, amount float64) {
    
}