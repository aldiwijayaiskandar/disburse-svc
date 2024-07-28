package user

import (
	"context"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func Consume(connection *amqp091.Connection){
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	exchange := "user.exchange"

	err = channel.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	RegisterGetQueue(exchange, channel)
}

func RegisterGetQueue(exchange string, channel *amqp091.Channel){
	// get user 
	getQueue, err := channel.QueueDeclare("user.get.queue", true, false, true, false, nil)
	if err != nil {
		panic(err)
	}
	
	err = channel.QueueBind(getQueue.Name, "user.get", exchange, false, nil)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	disburseConsumer, err := channel.ConsumeWithContext(ctx,getQueue.Name, "user-consumer", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	for messages := range disburseConsumer {
		log.Println(string(messages.Body))
	}
}