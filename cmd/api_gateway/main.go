package main

import (
	"github.com/gin-gonic/gin"
	"github.com/paper-assessment/internal/api_gateway/routes"
	"github.com/paper-assessment/pkg/config"
	"github.com/paper-assessment/pkg/rabbitmq"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		panic(err)
	}

	conn, err := rabbitmq.NewConnection(&cfg)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	r := gin.Default()

	routes.RegisterRoutes(r, conn)

	r.Run(cfg.Port)
}
