package main

import (
	"github.com/kennnyz/lamoda/internal/repository/postgres_repository"
	grpc_server "github.com/kennnyz/lamoda/internal/server/rpc_server"
	"github.com/kennnyz/lamoda/internal/service"
	"github.com/kennnyz/lamoda/pkg/database/postgres"
	"log"
)

func main() {

	db, err := postgres.NewClient("host=host.docker.internal port=5432 user=postgres password=password dbname=lamoda sslmode=disable timezone=UTC connect_timeout=5") // todo config
	if err != nil {
		log.Println(err)
		return
	}

	repo := postgres_repository.NewRepositories(db)

	services := service.NewServices(repo.Warehouse)
	log.Println(services.WareHouse)

	server := grpc_server.NewRPCServer(services, "0.0.0.0:12345")

	err = server.Run()
	if err != nil {
		log.Println(err)
	}
}
