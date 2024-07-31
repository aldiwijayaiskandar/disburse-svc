package main

import (
	"github.com/paper-assessment/internal/user/database"
	rabbitmq_delivery "github.com/paper-assessment/internal/user/delivery/rabbitmq"
	"github.com/paper-assessment/internal/user/domain/usecase"
	repository "github.com/paper-assessment/internal/user/repository/user"
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

	db := database.NewDatabaseConn(cfg.UserDatabaseUrl)

	userRepo := repository.NewUserRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepo)

	rabbitmq_delivery.Consume(conn, userUsecase)
}
