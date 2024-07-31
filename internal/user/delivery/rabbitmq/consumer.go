package rabbitmq_delivery

import (
	"github.com/streadway/amqp"
)

func Consume(conn *amqp.Connection) {
	// consumer, err := rabbitmq.NewConsumer(conn)
	// publisher, err := rabbitmq.NewPublisher(conn)

	// handler := handler.NewRabbitMQHandler(publisher)

	// if err != nil {
	// 	log.Fatalln("create consumer error")
	// }

	// listen to get user balance request
	forever := make(chan bool)
	go func() {

	}()

	<-forever
}
