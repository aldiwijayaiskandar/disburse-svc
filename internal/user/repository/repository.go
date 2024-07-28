package repository

import (
	"github.com/paper-assessment/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetById(id string) (*models.User){
	var user models.User

	if result := u.db.Where(&models.User{Id: id}).First(&user); result.Error == nil {
		return nil
	}

	return &user
}