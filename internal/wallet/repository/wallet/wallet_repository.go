package repository

import (
	repository_interface "github.com/paper-assessment/internal/wallet/domain/repository"
	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) repository_interface.WalletRepositoryInterface {
	return &WalletRepository{
		db: db,
	}
}
