package usecase

import (
	"github.com/rabbitmq/amqp091-go"
)

type UserUserCase struct {
	channel *amqp091.Channel
}

func NewUserUseCase(channel *amqp091.Channel) *UserUserCase {
	return &UserUserCase{
		channel: channel,
	}
}