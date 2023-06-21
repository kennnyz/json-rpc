package rpc_delivery

import (
	"github.com/gin-gonic/gin"
	"log"
)

type WarehouseProductsReq struct {
	WarehouseID  int   `json:"warehouse_id"`
	ProductCodes []int `json:"product_codes"`
}

type warehouseRemainingReq struct {
	WarehouseID int `json:"warehouse_id"`
}

type Reply struct {
	Data string
}

func (h *Handler) reserveProducts(c *gin.Context) {
	var warehouseProducts WarehouseProductsReq
	if err := c.ShouldBindJSON(&warehouseProducts); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Println(warehouseProducts)
	var reply Reply

	err := h.jsonRPC.Call("API.ReserveProducts", warehouseProducts, &reply)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(warehouseProducts)

	log.Println(reply)

	c.JSON(200, gin.H{"data": "ok"})
}

func (h *Handler) releaseProducts(c *gin.Context) {
	var warehouseProducts WarehouseProductsReq
	if err := c.ShouldBindJSON(&warehouseProducts); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Println(warehouseProducts)
	var reply Reply
	err := h.jsonRPC.Call("ReleaseProducts", warehouseProducts, reply)
	if err != nil {
		return
	}

	c.JSON(200, gin.H{"data": "ok"})
}

func (h *Handler) getRemainingProductCount(c *gin.Context) {
	var warehouseRemaining warehouseRemainingReq
	if err := c.ShouldBindJSON(&warehouseRemaining); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	reply := Reply{}
	err := h.jsonRPC.Call("GetRemainingProductCount", warehouseRemaining, reply)
	if err != nil {
		return
	}

	c.JSON(200, gin.H{"data": "ok"})

}
