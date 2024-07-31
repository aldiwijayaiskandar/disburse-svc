package handler

import (
	"encoding/json"

	"github.com/paper-assessment/internal/user/domain/models"
	constants "github.com/paper-assessment/pkg/contants"
	"github.com/paper-assessment/pkg/utils"
	"github.com/streadway/amqp"
)

func (h *RabbitMQHandler) GetUserHandler(d *amqp.Delivery) {
	var req models.GetUserRequest
	json.Unmarshal(d.Body, &req)

	// validate body
	err := utils.ValidateBody(req)

	if err != nil {
		errorMessage := err.Error()

		body, _ := json.Marshal(&models.GetUserResponse{
			Status:    constants.Error,
			User:      nil,
			ErrorCode: constants.InvalidRequest,
			Message:   &errorMessage,
		})

		h.publisher.Push(d.ReplyTo, body)
		return
	}

	res := h.userUsecase.GetUser(req)
	body, _ := json.Marshal(res)

	h.publisher.Push(d.ReplyTo, body)
}
