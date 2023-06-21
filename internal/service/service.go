package service

import "github.com/kennnyz/lamoda/lamodaTestTask/internal/models"

type WareHouse interface {
	ReserveProducts(wareHouseID int, productCodes int) error
	ReleaseReservedProducts(warehouseId, productCodes int) error
	GetRemainingProductCount(warehouseID int) ([]models.Product, error)
}

type Services struct {
	WareHouse WareHouse
}

func NewServices(wareHouse WareHouse) *Services {
	wareHouseService := NewWarehouseService(wareHouse)
	return &Services{
		WareHouse: wareHouseService,
	}
}
