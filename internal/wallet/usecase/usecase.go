package usecase

import (
	"github.com/paper-assessment/internal/wallet/repository"
	"github.com/rabbitmq/amqp091-go"
)

type UserUserCase struct {
	channel *amqp091.Channel
	repo *repository.WalletRepository
}

func NewUserUseCase(channel *amqp091.Channel, repo *repository.WalletRepository) *UserUserCase {
	return &UserUserCase{
		channel: channel,
		repo: repo,
	}
}