package repository

import (
	"errors"

	"github.com/paper-assessment/internal/user/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetById(id string) (*models.User, error) {
    var user models.User

    if result := u.db.Where(&models.User{Id: id}).First(&user); result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, result.Error
    }

    return &user, nil
}