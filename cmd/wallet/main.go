package main

import (
	"github.com/paper-assessment/internal/user/database"
	rabbitmq_delivery "github.com/paper-assessment/internal/wallet/delivery/rabbitmq"
	"github.com/paper-assessment/internal/wallet/domain/usecase"
	repository "github.com/paper-assessment/internal/wallet/repository/wallet"
	"github.com/paper-assessment/pkg/config"
	"github.com/paper-assessment/pkg/rabbitmq"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		panic(err)
	}

	conn, err := rabbitmq.NewConnection(&cfg)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	db := database.NewDatabaseConn(cfg.WalletDatabaseUrl)

	walletRepo := repository.NewWalletRepository(db)

	walletUsecase := usecase.NewWalletUsecase(walletRepo)

	rabbitmq_delivery.Consume(conn, walletUsecase)
}
