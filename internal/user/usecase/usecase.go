package usecase

import (
	"github.com/paper-assessment/internal/user/repository"
	"github.com/rabbitmq/amqp091-go"
)

type UserUserCase struct {
	channel *amqp091.Channel
	repo *repository.UserRepository
}

func NewUserUseCase(channel *amqp091.Channel, repo *repository.UserRepository) *UserUserCase {
	return &UserUserCase{
		channel: channel,
		repo: repo,
	}
}