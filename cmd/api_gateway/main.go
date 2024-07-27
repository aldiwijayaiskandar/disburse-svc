package main

import (
	"log"

	"github.com/gin-gonic/gin"
	apigateway "github.com/paper-assessment/internal/api_gateway"
	"github.com/paper-assessment/pkg/config"
)

func main(){
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	apigateway.RegisterRoutes(r)

	r.Run(c.Port)
}