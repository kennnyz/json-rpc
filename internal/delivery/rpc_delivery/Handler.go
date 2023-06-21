package rpc_delivery

import "github.com/kennnyz/lamoda/lamodaTestTask/internal/service"

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}
