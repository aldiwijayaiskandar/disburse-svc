package rabbitmq

import (
	"github.com/paper-assessment/pkg/config"
	"github.com/streadway/amqp"
)

// Initialize new RabbitMQ connection
func NewConnection(cfg *config.Config) (*amqp.Connection, error) {
	return amqp.Dial(cfg.BrokerUrl)
}

func getExchangeName() string {
	return "wallet_exchange"
}

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		getExchangeName(), // name
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
}
