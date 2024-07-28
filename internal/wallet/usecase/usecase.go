package usecase

import (
	"github.com/paper-assessment/internal/wallet/repository"
	"github.com/rabbitmq/amqp091-go"
)

type WalletUseCase struct {
	channel *amqp091.Channel
	repo *repository.WalletRepository
}

func NewWalletUseCase(channel *amqp091.Channel, repo *repository.WalletRepository) *WalletUseCase {
	return &WalletUseCase{
		channel: channel,
		repo: repo,
	}
}

