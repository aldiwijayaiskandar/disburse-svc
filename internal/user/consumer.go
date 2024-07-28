package user

import (
	"context"
	"encoding/json"
	"log"

	"github.com/paper-assessment/internal/models"
	"github.com/paper-assessment/internal/user/usecase"
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

	usecase := usecase.NewUserUseCase()

	RegisterGetUserQueue(exchange, channel, usecase)
}

func RegisterGetUserQueue(exchange string, channel *amqp091.Channel, usecase *usecase.UserUserCase){
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

		res :=  usecase.GetUser(models.GetUserRequest{
			Id: "test",
		})
	
		responseBytes, _ := json.Marshal(res)
		err = channel.Publish(
			"",
			messages.ReplyTo,
			false,
			false,
			amqp091.Publishing{
				ContentType:   "application/json",
				CorrelationId: messages.CorrelationId,
				Body:          responseBytes,
			},
		)

		if err != nil {
			log.Printf("Failed to publish a message: %v", err)
		}
	}
}