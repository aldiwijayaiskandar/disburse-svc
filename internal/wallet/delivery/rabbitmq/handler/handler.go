package handler

import (
	"github.com/paper-assessment/internal/wallet/domain/usecase"
	"github.com/paper-assessment/pkg/rabbitmq"
)

type RabbitMQHandler struct {
	publisher     rabbitmq.PublisherInterface
	walletUsecase usecase.WalletUsecaseInterface
}

func NewRabbitMQHandler(publisher rabbitmq.PublisherInterface, walletUsecase usecase.WalletUsecaseInterface) *RabbitMQHandler {
	return &RabbitMQHandler{
		publisher:     publisher,
		walletUsecase: walletUsecase,
	}
}
