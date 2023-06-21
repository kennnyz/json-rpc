package app

import (
	"github.com/kennnyz/lamoda/internal/delivery/rpc_delivery"
	"github.com/kennnyz/lamoda/internal/server/http_server"
	"net/rpc/jsonrpc"

	//"github.com/ybbus/jsonrpc/v3"
	"log"
)

func Run(configPath string) {
	// Запустить сервер

	//rpcServer := grpc_server.NewRPCServer("0.0.0.0:12345") // todo config
	//
	//go func() {
	//	err := rpcServer.Run()
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//}()

	// Слушаем по апи => rpc
	//rpcClient := jsonrpc.NewClient("localhost:12345")

	client, err := jsonrpc.Dial("tcp", "localhost:12345") //Меняется только эта строчка
	if err != nil {
		log.Fatal(err)
	}
	handler := rpc_delivery.NewHandler(client)

	httpServer := http_server.NewHTTPServer(":8080", handler.Init())

	err = httpServer.Run()
	if err != nil {
		log.Println(err)
		return
	}

}
