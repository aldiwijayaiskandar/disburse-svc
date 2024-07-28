package disburse

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/paper-assessment/internal/models"
	"github.com/rabbitmq/amqp091-go"
)

func Consume(connection *amqp091.Connection){
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	exchange := "disburse.exchange"

	err = channel.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	RegisterDisburseInitiateQueue(exchange, channel)
}

func RegisterDisburseInitiateQueue(exchange string, channel *amqp091.Channel){
	disburseQueue, err := channel.QueueDeclare("disburse.initiate.queue", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	
	err = channel.QueueBind(disburseQueue.Name, "disburse.initiate", exchange, false, nil)
	if err != nil {
		panic(err)
	}

	err = channel.QueueBind(disburseQueue.Name, "disburse.get.user.reply", exchange, false, nil)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	disburseConsumer, err := channel.ConsumeWithContext(ctx,disburseQueue.Name, "disburse-consumer", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	for messages := range disburseConsumer {
		var request models.DisburseRequest
		json.Unmarshal(messages.Body, &request)
		correlationId:= uuid.New().String()

		userRequestBytes, _ := json.Marshal(&models.GetUserRequest{
			Id: request.UserId,
		})
		// checking user exist
		err = channel.Publish(
			"user.exchange", // exchange
			"user.get.reply",      // routing key
			false,           // mandatory
			false,           // immediate
			amqp091.Publishing{
				ContentType:   "application/json",
				CorrelationId: correlationId,
				ReplyTo:       "user.get.reply.queue",
				Body:          userRequestBytes,
			},
		)

		if err != nil {
			panic(err)
		}

		ctx := context.Background()
		msgs, err := channel.ConsumeWithContext(ctx, "user.get.reply.queue" , "", true, false, false, false, nil)
		if err != nil {
			panic(err)
		}

		for d:= range msgs {
			if(correlationId == d.CorrelationId){
				var response models.GetUserResponse
				json.Unmarshal(d.Body, &response)
				log.Println("get reply: ", response)
			}
		}

		// get balance
	}
}