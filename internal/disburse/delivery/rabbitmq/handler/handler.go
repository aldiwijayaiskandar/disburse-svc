package handler

import (
	"github.com/paper-assessment/internal/disburse/domain/usecase"
	"github.com/paper-assessment/pkg/rabbitmq"
)

type RabbitMQHandler struct {
	publisher       rabbitmq.PublisherInterface
	disburseUsecase usecase.DisburseUsecaseInterface
}

func NewRabbitMQHandler(publisher rabbitmq.PublisherInterface, disburseUsecase usecase.DisburseUsecaseInterface) *RabbitMQHandler {
	return &RabbitMQHandler{
		publisher:       publisher,
		disburseUsecase: disburseUsecase,
	}
}
