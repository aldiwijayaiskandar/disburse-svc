package main

import (
	"log"

	"github.com/gin-gonic/gin"
	apigateway "github.com/paper-assessment/internal/api_gateway"
	"github.com/paper-assessment/pkg/config"
	"github.com/paper-assessment/pkg/rabbitmq"
)

func main(){
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	connection, err := rabbitmq.NewRabbitMQConn(&cfg)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	routes := gin.Default()

	apigateway.RegisterRoutes(routes, connection)

	routes.Run(cfg.Port)
}