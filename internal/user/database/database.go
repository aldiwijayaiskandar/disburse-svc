package database

import (
	"log"

	"github.com/paper-assessment/internal/user/database/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConn(url string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.DisableForeignKeyConstraintWhenMigrating = true
	db.AutoMigrate(&schema.User{})

	db.Save(&schema.User{
		Id:   "12345",
		Name: "test-user",
	})

	return db

}
