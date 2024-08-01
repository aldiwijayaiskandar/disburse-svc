package handler

import (
	"encoding/json"

	"github.com/paper-assessment/internal/disburse/domain/models"
	constants "github.com/paper-assessment/pkg/contants"
	"github.com/paper-assessment/pkg/utils"
	"github.com/streadway/amqp"
)

func (h *RabbitMQHandler) DisburseRequestHandler(d *amqp.Delivery) {
	var req models.DisburseRequest
	json.Unmarshal(d.Body, &req)

	// validate body
	err := utils.ValidateBody(req)

	if err != nil {
		errorMessage := err.Error()

		body, _ := json.Marshal(&models.DisburseResponse{
			Status:    constants.Error,
			ErrorCode: constants.InvalidRequest,
			Message:   &errorMessage,
		})

		h.publisher.Push(d.ReplyTo, body, d.CorrelationId)
		return
	}

	res := h.disburseUsecase.Disburse(req, d.CorrelationId)
	body, _ := json.Marshal(res)

	h.publisher.Push(d.ReplyTo, body, d.CorrelationId)
}
