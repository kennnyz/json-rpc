package rpc_delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/kennnyz/lamoda/internal/models"
	"log"
)

type WarehouseProductsReq struct {
	WarehouseID  int   `json:"warehouse_id"`
	ProductCodes []int `json:"product_codes"`
}

type Reply struct {
	Data string
}

type ReserveProductResponse struct {
	ReservedProductCodes []int
}

type ReleaseProductRequest struct {
	WarehouseID  int   `json:"warehouse_id"`
	ProductCodes []int `json:"product_codes"`
}

type ReleaseProductResponse struct {
	ReleasedProductCodes []int
}

type GetRemainingProductCountRequest struct {
	WarehouseID int `json:"warehouse_id"`
}

type GetRemainingProductCountResponse struct {
	Products []models.Product
}

func (h *Handler) reserveProducts(c *gin.Context) {
	var warehouseProducts WarehouseProductsReq
	if err := c.ShouldBindJSON(&warehouseProducts); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var reply ReserveProductResponse

	err := h.jsonRPC.Call("API.ReserveProducts", warehouseProducts, &reply)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"Reserved": reply.ReservedProductCodes})
}

func (h *Handler) releaseProducts(c *gin.Context) {
	var releaseProductRequest ReleaseProductRequest
	if err := c.ShouldBindJSON(&releaseProductRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var reply ReleaseProductResponse
	err := h.jsonRPC.Call("API.ReleaseProducts", releaseProductRequest, &reply)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{"Released Products ": reply.ReleasedProductCodes})
}

func (h *Handler) getRemainingProductCount(c *gin.Context) {
	var warehouseRemaining GetRemainingProductCountRequest
	if err := c.ShouldBindJSON(&warehouseRemaining); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var reply GetRemainingProductCountResponse
	err := h.jsonRPC.Call("API.GetRemainingProductCount", warehouseRemaining, &reply)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{"Remaining Products ": reply.Products})

}
