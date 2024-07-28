package wallet

import (
	// "github.com/paper-assessment/internal/user/repository"
	// "github.com/paper-assessment/internal/user/usecase"
	"github.com/paper-assessment/internal/wallet/repository"
	"github.com/rabbitmq/amqp091-go"
)

func Consume(connection *amqp091.Connection, repo *repository.WalletRepository){
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	exchange := "user.exchange"

	err = channel.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	// _ := usecase.NewUserUseCase(
	// 	channel,
	// 	repo,
	// )
}