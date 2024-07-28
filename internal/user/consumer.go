package user

import (
	"context"
	"encoding/json"
	"log"

	"github.com/paper-assessment/internal/models"
	"github.com/paper-assessment/internal/user/repository"
	"github.com/paper-assessment/internal/user/usecase"
	"github.com/rabbitmq/amqp091-go"
)

func Consume(connection *amqp091.Connection, repo *repository.UserRepository){
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	exchange := "user.exchange"

	err = channel.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	getUserReplyQueue, err := channel.QueueDeclare("user.get.reply.queue", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	err = channel.QueueBind(getUserReplyQueue.Name ,"user.get.reply",exchange, false, nil)

	if  err != nil {
		panic(err)
	}

	usecase := usecase.NewUserUseCase(
		channel,
		repo,
	)

	RegisterGetUserQueue(exchange, channel, usecase)
}

func RegisterGetUserQueue(exchange string, channel *amqp091.Channel, usecase *usecase.UserUserCase){
	getQueue, err := channel.QueueDeclare("user.get.queue", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	
	err = channel.QueueBind(getQueue.Name, "user.get", exchange, false, nil)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	msgs, err := channel.ConsumeWithContext(ctx,getQueue.Name, "user-consumer", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	for d := range msgs {
		log.Println("correlation id: ",d.CorrelationId)

		var request models.GetUserRequest
		json.Unmarshal(d.Body, &request)

		res := usecase.GetUser(models.GetUserRequest{
			Id: request.Id,
		})

		log.Println("response: ", res)

		responseBytes, _ := json.Marshal(res)

		channel.Publish(
			exchange,
			d.ReplyTo,
			false,
			false,
			amqp091.Publishing{
				ContentType:   "application/json",
				CorrelationId: d.CorrelationId,
				Body:          responseBytes,
			},
		)
	}
}