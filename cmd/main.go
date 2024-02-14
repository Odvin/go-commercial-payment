package main

import (
	"log"

	"github.com/Odvin/go-commercial-payment/internal/adapters/db"
	"github.com/Odvin/go-commercial-payment/internal/adapters/grpc"
	"github.com/Odvin/go-commercial-payment/internal/application/core/api"
	"github.com/huseyinbabal/microservices/payment/config"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())

	grpcAdapter.Run()
}
