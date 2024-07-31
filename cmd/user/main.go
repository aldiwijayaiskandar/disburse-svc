package main

import (
	"github.com/paper-assessment/internal/user/database"
	rabbitmq_delivery "github.com/paper-assessment/internal/user/delivery/rabbitmq"
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

	database.NewDatabaseConn(cfg.UserDatabaseUrl)

	rabbitmq_delivery.Consume(conn)
}
