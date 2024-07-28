package rabbitmq

import (
	"github.com/paper-assessment/pkg/config"
	"github.com/rabbitmq/amqp091-go"
)

// Initialize new RabbitMQ connection
func NewRabbitMQConn(cfg *config.Config) (*amqp091.Connection, error) {
	return amqp091.Dial(cfg.BrokerUrl)
}