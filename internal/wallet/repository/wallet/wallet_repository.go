package repository

import (
	"github.com/paper-assessment/internal/wallet/domain/repository"
	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) repository.WalletRepository {
	return &WalletRepository{
		db: db,
	}
}
