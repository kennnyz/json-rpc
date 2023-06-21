package rpc_delivery

import (
	"github.com/gin-gonic/gin"
	"net/rpc"
)

type Handler struct {
	jsonRPC *rpc.Client

	// клиент json-rpc
}

func NewHandler(rpcClient *rpc.Client) *Handler {
	return &Handler{
		jsonRPC: rpcClient,
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
