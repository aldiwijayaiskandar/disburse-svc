package wallet

import (
	"github.com/paper-assessment/internal/wallet/delivery/rabbitmq"
	"github.com/paper-assessment/internal/wallet/repository"
	"github.com/paper-assessment/internal/wallet/usecase"
	"github.com/rabbitmq/amqp091-go"
)

func Consume(connection *amqp091.Connection, repo *repository.WalletRepository){
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	exchange := "wallet.exchange"

	err = channel.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	usecase := usecase.NewWalletUseCase(
		channel,
		repo,
	)

	rabbitmq.RegisterGetUserBalanceQueue(channel, usecase)
	rabbitmq.RegisterDeductUserBalanceQueue(channel, usecase)
}