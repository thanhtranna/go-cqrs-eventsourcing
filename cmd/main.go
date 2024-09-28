package main

import (
	"flag"
	"log"

	"github.com/thanhtranna/go-cqrs-eventsourcing/config"
	"github.com/thanhtranna/go-cqrs-eventsourcing/internal/server"
	"github.com/thanhtranna/go-cqrs-eventsourcing/pkg/logger"
)

// @contact.name Thanh Tran
// @contact.url https://github.com/thanhtranna
// @contact.email tranthanh19589@gmail.com
func main() {
	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	appLogger.WithName(server.GetMicroserviceName(cfg))
	appLogger.Fatal(server.NewServer(cfg, appLogger).Run())
}
