package database

import (
	"github.com/paper-assessment/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func NewDatabaseConn(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("can't connect to database")
	}

	db.AutoMigrate(&models.User{})
	return db
}