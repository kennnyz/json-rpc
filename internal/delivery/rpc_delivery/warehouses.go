package rpc_delivery

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) initWarehousesRoutes(api *gin.RouterGroup) {
	warehouses := api.Group("/warehouses")
	{
		warehouses.POST("/reserve", h.reserveProducts)
		warehouses.POST("/release", h.releaseProducts)
		warehouses.GET("/remaining", h.getRemainingProductCount)
	}
}

type warehouseProductsReq struct {
	WarehouseID  int   `json:"warehouse_id"`
	ProductCodes []int `json:"product_codes"`
}

type warehouseRemainingReq struct {
	WarehouseID int `json:"warehouse_id"`
}

func (h *Handler) reserveProducts(c *gin.Context) {
	var warehouseProducts warehouseProductsReq
	if err := c.ShouldBindJSON(&warehouseProducts); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	call, err := h.jsonRPC.Call(c, "ReserveProducts", warehouseProducts)
	if err != nil {
		return
	}

	c.JSON(200, call)
}

func (h *Handler) releaseProducts(c *gin.Context) {
	var warehouseProducts warehouseProductsReq
	if err := c.ShouldBindJSON(&warehouseProducts); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	call, err := h.jsonRPC.Call(c, "ReleaseProducts", warehouseProducts)
	if err != nil {
		return
	}

	c.JSON(200, call)
}

func (h *Handler) getRemainingProductCount(c *gin.Context) {
	var warehouseRemaining warehouseRemainingReq
	if err := c.ShouldBindJSON(&warehouseRemaining); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	call, err := h.jsonRPC.Call(c, "GetRemainingProductCount", warehouseRemaining)
	if err != nil {
		return
	}

	c.JSON(200, call)

}
