package repository

import (
	repository_interface "github.com/paper-assessment/internal/user/domain/repository"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository_interface.UserRepositoryInterface {
	return &UserRepository{
		db: db,
	}
}
