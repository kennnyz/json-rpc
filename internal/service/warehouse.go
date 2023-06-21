package service

import (
	"github.com/kennnyz/lamoda/internal/models"
	"github.com/kennnyz/lamoda/internal/repository/postgres_repository"
)

type WarehouseService struct {
	repo postgres_repository.Warehouse
}

func NewWarehouseService(repo postgres_repository.Warehouse) *WarehouseService {
	return &WarehouseService{repo: repo}
}

func (s *WarehouseService) ReserveProducts(wareHouseID int, productCodes int) error {
	return s.repo.ReserveProducts(wareHouseID, productCodes)
}

func (s *WarehouseService) ReleaseReservedProducts(warehouseId, productCodes int) error {
	return s.repo.ReleaseReservedProducts(warehouseId, productCodes)
}

func (s *WarehouseService) GetRemainingProductCount(warehouseID int) ([]models.Product, error) {
	return s.repo.GetRemainingProductCount(warehouseID)
}
