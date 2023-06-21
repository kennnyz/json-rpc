package app

import (
	"github.com/kennnyz/lamoda/lamodaTestTask/internal/delivery/rpc_delivery"
	"github.com/kennnyz/lamoda/lamodaTestTask/internal/repository/postgres_repository"
	"github.com/kennnyz/lamoda/lamodaTestTask/internal/server/http_server"
	grpc_server "github.com/kennnyz/lamoda/lamodaTestTask/internal/server/rpc_server"
	"github.com/kennnyz/lamoda/lamodaTestTask/internal/service"
	"github.com/kennnyz/lamoda/lamodaTestTask/pkg/database/postgres"
	"github.com/ybbus/jsonrpc/v3"
	"log"
)

func Run(configPath string) {
	// Запустить сервер

	rpcServer := grpc_server.NewRPCServer("0.0.0.0:12345") // todo config

	go func() {
		err := rpcServer.Run()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	db, err := postgres.NewClient("host=localhost port=5432 user=postgres password=password dbname=lamoda sslmode=disable timezone=UTC connect_timeout=5") // todo config
	if err != nil {
		log.Println(err)
		return
	}

	repo := postgres_repository.NewRepositories(db)

	services := service.NewServices(repo.Warehouse)
	// Слушаем по апи => rpc
	rpcClient := jsonrpc.NewClient("0.0.0.0:12345")
	handler := rpc_delivery.NewHandler(services, rpcClient)

	httpServer := http_server.NewHTTPServer(":8080", handler.Init())

	go func() {
		err := httpServer.Run()
		if err != nil {
			log.Println(err)
			return
		}
	}()

}
