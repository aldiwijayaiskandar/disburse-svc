package main

import (
	"github.com/paper-assessment/internal/disburse"
	"github.com/paper-assessment/pkg/config"
	"github.com/paper-assessment/pkg/rabbitmq"
)

func main(){
	cfg, err := config.LoadConfig()

	if err != nil {
		panic(err)
	}

	connection, err := rabbitmq.NewRabbitMQConn(&cfg)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	disburse.Consume(connection)
}