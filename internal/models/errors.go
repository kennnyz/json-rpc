package models

import "errors"

var (
	ErrNoProducts          = errors.New("no products in warehouse")
	ErrAssortmentNotExists = errors.New("assortment not exists")
	ErrNoProductsInReserve = errors.New("no product in reserve")
)
