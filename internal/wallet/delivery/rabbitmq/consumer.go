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
	publisher, err := rabbitmq.NewPublisher(conn)

	handler := handler.NewRabbitMQHandler(publisher, usecase)

	if err != nil {
		log.Fatalln("create consumer error")
	}

	// listen to get user balance request
	forever := make(chan bool)
	go func() {
		go consumer.Listen(
			[]string{"wallet.balance.get.request"},
			func(delivery *amqp.Delivery) {
				handler.GetUserBalanceHandler(delivery)
			},
		)

		go consumer.Listen(
			[]string{"wallet.balance.deduct.request"},
			func(delivery *amqp.Delivery) {
				handler.DeductBalanceHandler(delivery)
			},
		)
	}()

	<-forever
}
