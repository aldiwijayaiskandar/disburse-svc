package main

import (
	"github.com/paper-assessment/internal/user/database"
	rabbitmq_delivery "github.com/paper-assessment/internal/user/delivery/rabbitmq"
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

	repository.NewUserRepository(db)

	rabbitmq_delivery.Consume(conn)
}
