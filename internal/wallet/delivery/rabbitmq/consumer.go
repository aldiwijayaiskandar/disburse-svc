package rabbitmq_delivery

import (
	"log"

	"github.com/paper-assessment/internal/wallet/delivery/rabbitmq/handler"
	"github.com/paper-assessment/internal/wallet/domain/usecase"
	"github.com/paper-assessment/pkg/rabbitmq"
	"github.com/streadway/amqp"
)

func Consume(conn *amqp.Connection, usecase usecase.WalletUsecaseInterface) {
	consumer, err := rabbitmq.NewConsumer(conn)

	if err != nil {
		log.Fatalln("create consumer error")
	}

	// listen to get user balance request
	consumer.Listen(
		[]string{"wallet.balance.get"},
		func(delivery *amqp.Delivery) {
			handler.GetUserBalanceHandler(delivery, usecase)
		},
	)
}
