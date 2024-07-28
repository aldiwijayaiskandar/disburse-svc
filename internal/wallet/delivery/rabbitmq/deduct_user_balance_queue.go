package rabbitmq

import (
	"context"
	"encoding/json"

	"github.com/paper-assessment/internal/models"
	"github.com/paper-assessment/internal/wallet/usecase"
	"github.com/rabbitmq/amqp091-go"
)

func RegisterDeductUserBalanceQueue(channel *amqp091.Channel, usecase *usecase.WalletUseCase){
	// declaring all queue for get balance
	_, err := channel.QueueDeclare("wallet.deduct.balance.queue", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	err = channel.QueueBind("wallet.deduct.balance.queue", "wallet.deduct.balance", "wallet.exchange", false, nil)

	if err != nil {
		panic(err)
	}

	ctx :=context.Background()
	msgs, err := channel.ConsumeWithContext(ctx, "wallet.deduct.balance.queue", "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	for d := range msgs {
		var request models.DeductUserBalanceRequest
		json.Unmarshal(d.Body, &request)

		usecase.DeductUserBalance(request)
	}
}