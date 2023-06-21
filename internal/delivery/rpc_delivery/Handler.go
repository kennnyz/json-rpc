package rpc_delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/kennnyz/lamoda/lamodaTestTask/internal/service"
	"github.com/ybbus/jsonrpc/v3"
)

// RPC server

type Handler struct {
	services *service.Services
	jsonRPC  jsonrpc.RPCClient

	// клиент json-rpc
}

func NewHandler(services *service.Services, rpcClient jsonrpc.RPCClient) *Handler {
	return &Handler{
		services: services,
		jsonRPC:  rpcClient,
	}
}

// Init gin router

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	router.POST("/reserve", h.reserveProducts)
	router.POST("/release", h.releaseProducts)
	router.GET("/remaining", h.getRemainingProductCount)

	return router
}
