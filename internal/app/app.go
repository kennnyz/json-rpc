package app

import (
	"github.com/kennnyz/lamoda/internal/delivery/rpc"
	"github.com/kennnyz/lamoda/internal/server/http_server"
	"net/rpc/jsonrpc"
	"os"

	//"github.com/ybbus/jsonrpc/v3"
	"log"
)

func Run(configPath string) {
	tcpAddr := os.Getenv("RCP_ADDRESS")
	RpcClient, err := jsonrpc.Dial("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	handler := rpc.NewHandler(RpcClient)

	httpAddr := os.Getenv("HTTP_ADDRESS")
	httpServer := http_server.NewHTTPServer(httpAddr, handler.Init())

	err = httpServer.Run()
	if err != nil {
		log.Println(err)
		return
	}
}
