package app

import (
	"github.com/kennnyz/lamoda/internal/delivery/rpc_delivery"
	"github.com/kennnyz/lamoda/internal/server/http_server"
	"net/rpc/jsonrpc"

	//"github.com/ybbus/jsonrpc/v3"
	"log"
)

func Run(configPath string) {
	RpcClient, err := jsonrpc.Dial("tcp", "localhost:12345") //Меняется только эта строчка
	if err != nil {
		log.Fatal(err)
	}
	handler := rpc_delivery.NewHandler(RpcClient)

	httpServer := http_server.NewHTTPServer(":8080", handler.Init())

	err = httpServer.Run()
	if err != nil {
		log.Println(err)
		return
	}

}
