package app

import (
	"github.com/kennnyz/lamoda/internal/delivery/rpc"
	"github.com/kennnyz/lamoda/internal/server/http_server"
	"net/rpc/jsonrpc"

	//"github.com/ybbus/jsonrpc/v3"
	"log"
)

func Run(configPath string) {
	RpcClient, err := jsonrpc.Dial("tcp", "localhost:12345") //TODO config
	if err != nil {
		log.Fatal(err)
	}
	handler := rpc.NewHandler(RpcClient)

	httpServer := http_server.NewHTTPServer(":8080", handler.Init()) // todo config

	err = httpServer.Run()
	if err != nil {
		log.Println(err)
		return
	}
}
