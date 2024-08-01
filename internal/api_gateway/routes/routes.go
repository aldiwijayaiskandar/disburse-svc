package routes

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/paper-assessment/internal/api_gateway/domain/models"
	disburseModel "github.com/paper-assessment/internal/disburse/domain/models"
	"github.com/paper-assessment/pkg/rabbitmq"
	"github.com/streadway/amqp"
)

func RegisterRoutes(r *gin.Engine, conn *amqp.Connection) {
	routes := r.Group("/v1/disburse")

	routes.POST("/", func(ctx *gin.Context) {
		var body disburseModel.DisburseRequest

		if err := ctx.BindJSON(&body); err != nil {
			errorMessage := "Invalid request body"
			ctx.JSON(400, models.ApiResponse{
				Status:  400,
				Message: &errorMessage,
			})

			return
		}

		correlationId := uuid.New().String()

		publisher, err := rabbitmq.NewPublisher(conn)

		if err != nil {
			log.Fatalln("fail to create publisher")
		}

		consumer, err := rabbitmq.NewConsumer(conn)

		if err != nil {
			log.Fatalln("fail to create consumer")
		}

		byteBody, _ := json.Marshal(body)
		publisher.Push("disburse.request", byteBody, correlationId)

		d, err := consumer.WaitReply(correlationId)

		log.Println(d.Body)

		var res disburseModel.DisburseResponse
		json.Unmarshal(d.Body, &res)

		if err != nil {
			errorMessage := "Internal server error"
			ctx.JSON(500, models.ApiResponse{
				Status:  500,
				Message: &errorMessage,
			})

			return
		}

		ctx.JSON(res.ErrorCode.StatusCode(), models.ApiResponse{
			Status: res.ErrorCode.StatusCode(),
			Data: map[string]interface{}{
				"balance": res.Balance,
			},
		})
	})
}
