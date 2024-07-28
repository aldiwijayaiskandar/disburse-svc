package apigateway

import (
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/paper-assessment/internal/models"
	"github.com/rabbitmq/amqp091-go"
)

func RegisterRoutes(routes *gin.Engine, connection *amqp091.Connection) *Client {
	client := &Client{
		connection: connection,
	}

	routes.POST("/disburse", client.Disburse)

	return client
}

func (client *Client) Disburse(ctx *gin.Context) {
	var body models.DisburseRequest
	if err := ctx.BindJSON(&body); err != nil {
		panic(err)
	}

	correlationId := uuid.New().String()

	channel, err := client.connection.Channel()
	if err != nil {
		panic(err)
	}
	
	// public to disburse service
	responseBytes, _ := json.Marshal(body)
	channel.Publish(
		"disburse.exchange",
		"disburse.initiate",
		false,           // mandatory
		false,           // immediate
		amqp091.Publishing{
			ContentType:   "application/json",
			CorrelationId: correlationId,
			Body:          responseBytes,
		},
	)

	// waiting for reply
	c, cancel := context.WithCancel(context. Background())
	msgs, err := channel.ConsumeWithContext(
		c,
		"disburse.initiate.reply.queue",
		"",
		true, 
		false, 
		false, 
		false, 
		nil,
	)

	if err != nil {
		panic(err)
	}

	for d := range msgs {
		if(d.CorrelationId == correlationId) {
			var response models.ApiResponse
			json.Unmarshal(d.Body, &response)

			ctx.JSON(int(response.Status), response)
			break
		}
	}

	cancel()
}