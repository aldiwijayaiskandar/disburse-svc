package rabbitmq_delivery

import (
	"github.com/paper-assessment/internal/disburse/data/services"
	"github.com/paper-assessment/internal/disburse/delivery/rabbitmq/handler"
	"github.com/paper-assessment/internal/disburse/domain/usecase"
	"github.com/paper-assessment/pkg/rabbitmq"
	"github.com/streadway/amqp"
)

func Consume(conn *amqp.Connection) {
	consumer, err := rabbitmq.NewConsumer(conn)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	publisher, err := rabbitmq.NewPublisher(conn)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	userService := services.NewUserService(consumer, publisher)
	wallerService := services.NewWalletService(consumer, publisher)

	disburseUsecase := usecase.NewDisburseUsecase(userService, wallerService)

	handler := handler.NewRabbitMQHandler(publisher, disburseUsecase)

	// listen to get user balance request
	forever := make(chan bool)
	go func() {
		consumer.Listen([]string{"disburse.request"}, handler.DisburseRequestHandler)
	}()

	<-forever
}
