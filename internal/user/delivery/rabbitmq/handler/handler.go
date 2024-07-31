package handler

import (
	"github.com/paper-assessment/internal/user/domain/usecase"
	"github.com/paper-assessment/pkg/rabbitmq"
)

type RabbitMQHandler struct {
	publisher   rabbitmq.PublisherInterface
	userUsecase usecase.UserUsecaseInterface
}

func NewRabbitMQHandler(publisher rabbitmq.PublisherInterface, userUsecase usecase.UserUsecaseInterface) *RabbitMQHandler {
	return &RabbitMQHandler{
		publisher:   publisher,
		userUsecase: userUsecase,
	}
}
