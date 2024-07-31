package rabbitmq_delivery

import (
	"log"

	"github.com/paper-assessment/internal/user/delivery/rabbitmq/handler"
	"github.com/paper-assessment/internal/user/domain/usecase"
	"github.com/paper-assessment/pkg/rabbitmq"
	"github.com/streadway/amqp"
)

func Consume(conn *amqp.Connection, userUsecase usecase.UserUsecaseInterface) {
	consumer, err := rabbitmq.NewConsumer(conn)
	publisher, err := rabbitmq.NewPublisher(conn)

	handler := handler.NewRabbitMQHandler(publisher, userUsecase)

	if err != nil {
		log.Fatalln("create consumer error")
	}

	// listen to get user balance request
	forever := make(chan bool)
	go func() {
		consumer.Listen([]string{"user.get.request"}, handler.GetUserHandler)
	}()

	<-forever
}
