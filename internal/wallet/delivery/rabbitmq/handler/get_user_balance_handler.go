package handler

import (
	"encoding/json"

	"github.com/paper-assessment/internal/wallet/domain/models"
	constants "github.com/paper-assessment/pkg/contants"
	"github.com/paper-assessment/pkg/utils"
	"github.com/streadway/amqp"
)

func (h *RabbitMQHandler) GetUserBalanceHandler(d *amqp.Delivery) {
	var req models.GetUserBalanceRequest
	json.Unmarshal(d.Body, &req)

	// validate body
	err := utils.ValidateBody(req)

	if err != nil {
		errorMessage := err.Error()

		body, _ := json.Marshal(&models.GetUserBalanceResponse{
			Status:    constants.Error,
			Balance:   nil,
			ErrorCode: constants.InvalidRequest,
			Message:   &errorMessage,
		})

		h.publisher.Push(d.ReplyTo, body, d.CorrelationId)
		return
	}

	res := h.walletUsecase.GetUserBalance(req.UserId)
	body, _ := json.Marshal(res)

	h.publisher.Push(d.ReplyTo, body, d.CorrelationId)
}
