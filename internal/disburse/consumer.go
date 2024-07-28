package disburse

import (
	"context"
	"encoding/json"

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

	disburseQueue, err := channel.QueueDeclare("disburse.initiate.queue", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	
	err = channel.QueueBind(disburseQueue.Name, "disburse.initiate", exchange, false, nil)
	if err != nil {
		panic(err)
	}

	// reply
	disburseReplyQueue, err := channel.QueueDeclare("disburse.initiate.reply.queue", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	err = channel.QueueBind(disburseReplyQueue.Name, "disburse.initiate.reply", exchange, false, nil)
	if err != nil {
		panic(err)
	}

	RegisterDisburseInitiateQueue(exchange, channel)
}

func RegisterDisburseInitiateQueue(exchange string, channel *amqp091.Channel){
	ctx := context.Background()
	disburseMsgs, err := channel.ConsumeWithContext(ctx, "disburse.initiate.queue", "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	for msgs := range disburseMsgs {
		var request models.DisburseRequest
		json.Unmarshal(msgs.Body, &request)

		checkUser(msgs.CorrelationId, channel, &models.GetUserRequest{
			Id: request.UserId,
		})

		// get balance
		
	}
}

func checkUser (correlationId string, channel *amqp091.Channel, request *models.GetUserRequest) {
	userRequestBytes, _ := json.Marshal(request)

	// checking user exist
	err := channel.Publish(
		"user.exchange", // exchange
		"user.get.reply", // routing key
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

	ctx, cancel := context.WithCancel(context.Background())

	getUserReplyMsgs, err := channel.ConsumeWithContext(ctx, "user.get.reply.queue" , "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	for userReplyMsgs:= range getUserReplyMsgs {
		if(correlationId == userReplyMsgs.CorrelationId){
			var response models.GetUserResponse
			json.Unmarshal(userReplyMsgs.Body, &response)
			
			if response.Error != nil {
				ReplyDisburseInitiate(
					channel,
					correlationId,
					&models.ApiResponse{
						Status: 400,
						Error: response.Error,
					},
				)
			}

			if response.Data == nil { 
				err := "User Doesn't Exist"

				ReplyDisburseInitiate(
					channel,
					correlationId,
					&models.ApiResponse{
						Status: 400,
						Error: &err,
					},
				)
			}
			
			break;
		}
	}

	cancel()
}

func ReplyDisburseInitiate(channel *amqp091.Channel, correlationId string, response *models.ApiResponse,){
	responseBytes, _ := json.Marshal(response)
	err := channel.Publish(
		"disburse.exchange", // exchange
		"disburse.initiate.reply",      // routing key
		false,           // mandatory
		false,           // immediate
		amqp091.Publishing{
			ContentType:   "application/json",
			CorrelationId: correlationId,
			Body:          responseBytes,
		},
	)

	if err != nil {
		panic(err)
	}
}