package handler

import (
	"encoding/json"

	"github.com/paper-assessment/internal/models"
	"github.com/paper-assessment/internal/wallet/domain/usecase"
	"github.com/streadway/amqp"
)

func GetUserBalanceHandler(d *amqp.Delivery, usecase usecase.WalletUsecaseInterface) {
	var req models.GetUserBalanceRequest
	err := json.Unmarshal(d.Body, &req)

	if err != nil {
		return
	}
}
