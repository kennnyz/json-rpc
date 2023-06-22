package main

import (
	"github.com/kennnyz/lamoda/internal/repository/postgres_repository"
	grpc_server "github.com/kennnyz/lamoda/internal/server/rpc_server"
	"github.com/kennnyz/lamoda/internal/service"
	"github.com/kennnyz/lamoda/pkg/database/postgres"
	"log"
	"os"
)

func main() {

	dsn := os.Getenv("DSN")
	db, err := postgres.NewClient(dsn)
	if err != nil {
		log.Println(err)
		return
	}

	repo := postgres_repository.NewRepositories(db)

	services := service.NewServices(repo.Warehouse)
	log.Println(services.WareHouse)

	rcpAddr := os.Getenv("RCP_ADDRESS")
	server := grpc_server.NewRPCServer(services, rcpAddr)

	err = server.Run()
	if err != nil {
		log.Println(err)
	}
}
