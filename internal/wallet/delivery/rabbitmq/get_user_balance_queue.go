package rabbitmq

import (
	"context"
	"encoding/json"

	"github.com/paper-assessment/internal/models"
	"github.com/paper-assessment/internal/wallet/usecase"
	"github.com/rabbitmq/amqp091-go"
)

func RegisterGetUserBalanceQueue(channel *amqp091.Channel, usecase *usecase.WalletUseCase){
	// declaring all queue for get balance
	_, err := channel.QueueDeclare("wallet.get.balance.queue", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	_, err = channel.QueueDeclare("wallet.get.balance.reply.queue", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	err = channel.QueueBind("wallet.get.balance.queue", "wallet.get.balance", "wallet.exchange", false, nil)

	if err != nil {
		panic(err)
	}

	err = channel.QueueBind("wallet.get.balance.reply.queue", "wallet.get.reply.balance", "wallet.exchange", false, nil)

	if err != nil {
		panic(err)
	}

	ctx :=context.Background()
	msgs, err := channel.ConsumeWithContext(ctx, "wallet.get.balance.queue", "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	for d := range msgs {
		var request models.GetUserBalanceRequest
		json.Unmarshal(d.Body, &request)

		res, err := usecase.GetUserBalance(request)

		if err != nil {
			err := err.Error()
			responseBytes, _ := json.Marshal(&models.GetUserBalanceResponse{
				Balance: res.Balance,
				Error: &err,
			})
			channel.Publish(
				"wallet.exchange",
				"wallet.get.balance.reply",
				false,
				false,
				amqp091.Publishing{
					ContentType:   "application/json",
					CorrelationId: d.CorrelationId,
					Body:          responseBytes,
				},
			)
			break
		}

		responseBytes, _ := json.Marshal(&models.GetUserBalanceResponse{
			Balance: res.Balance,
		})
		channel.Publish(
			"wallet.exchange",
			"wallet.get.balance.reply",
			false,
			false,
			amqp091.Publishing{
				ContentType:   "application/json",
				CorrelationId: d.CorrelationId,
				Body:          responseBytes,
			},
		)
	}
}