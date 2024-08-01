package database

import (
	"log"

	"github.com/paper-assessment/internal/wallet/database/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConn(url string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.DisableForeignKeyConstraintWhenMigrating = true
	db.AutoMigrate(&schema.Wallet{})

	db.Save(&schema.Wallet{
		UserId:  "12345",
		Balance: 50000.00,
	})

	return db

}
