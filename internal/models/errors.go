package models

import "errors"

var (
	ErrNoProducts            = errors.New("no products in warehouse")
	ErrNoProductsInReserve   = errors.New("no product in reserve")
	ErrWarehouseNotAvailable = errors.New("warehouse not available")
)
