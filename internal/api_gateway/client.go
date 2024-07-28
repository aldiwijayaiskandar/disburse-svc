package apigateway

import "github.com/rabbitmq/amqp091-go"

type Client struct {
	connection *amqp091.Connection
}
