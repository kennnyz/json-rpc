package app

import (
	"fmt"
	"github.com/kennnyz/lamoda/internal/delivery/rpc"
	"github.com/kennnyz/lamoda/internal/server/http_server"
	"net/rpc/jsonrpc"
	"os"
	"time"

	//"github.com/ybbus/jsonrpc/v3"
	"log"
)

func Run(configPath string) {
	tcpAddr := os.Getenv("RCP_ADDRESS")
	time.Sleep(30 * time.Second)
	RpcClient, err := jsonrpc.Dial("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ok rpc")
	handler := rpc.NewHandler(RpcClient)

	httpAddr := os.Getenv("HTTP_ADDRESS")
	httpServer := http_server.NewHTTPServer(httpAddr, handler.Init())

	fmt.Println("Server is listening..." + httpAddr)
	err = httpServer.Run()
	if err != nil {
		log.Println(err)
		return
	}
}
