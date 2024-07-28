package main

import (
	"github.com/paper-assessment/internal/user"
	"github.com/paper-assessment/internal/user/database"
	"github.com/paper-assessment/internal/user/repository"
	"github.com/paper-assessment/pkg/config"
	"github.com/paper-assessment/pkg/rabbitmq"
)

func main(){
	cfg, err := config.LoadConfig()

	if err != nil {
		panic(err)
	}

	connection, err := rabbitmq.NewRabbitMQConn(&cfg)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	db := database.NewDatabaseConn(cfg.WalletDatabaseUrl)

	repository := repository.NewUserRepository(db)

	user.Consume(connection, repository)
}