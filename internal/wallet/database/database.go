package database

import (
	"github.com/paper-assessment/internal/wallet/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func NewDatabaseConn(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic("can't connect to database")
	}

	db.AutoMigrate(&models.Wallet{})
	return db
}