package grpc_server

import (
	"github.com/kennnyz/lamoda/internal/models"
	"github.com/kennnyz/lamoda/internal/service"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Config struct {
	ListenAddr string
}

type Server struct {
	services *service.Services
	Config
}

func NewRPCServer(services *service.Services, addr string) *Server {
	return &Server{
		services: services,
		Config: Config{
			ListenAddr: addr,
		},
	}
}

type API struct {
	services *service.Services
}

type WarehouseProductsReq struct {
	WarehouseID  int   `json:"warehouse_id"`
	ProductCodes []int `json:"product_codes"`
}

type ReserveProductResponse struct {
	ReservedProductCodes []int
}

type ReleaseProductRequest struct {
	WarehouseID  int
	ProductCodes []int
}

type ReleaseProductResponse struct {
	ReleasedProductCodes []int
}

type GetRemainingProductCountRequest struct {
	WarehouseID int
}

type GetRemainingProductCountResponse struct {
	Products []models.Product
}

func (a *API) ReserveProducts(req *WarehouseProductsReq, res *ReserveProductResponse) error {
	var reserved []int
	log.Println(req.ProductCodes)
	for _, productCode := range req.ProductCodes {
		err := a.services.WareHouse.ReserveProducts(req.WarehouseID, productCode)
		if err != nil && err != models.ErrNoProducts {
			return err
		}
		reserved = append(reserved, productCode)
	}
	*res = ReserveProductResponse{
		ReservedProductCodes: reserved,
	}
	log.Println("hello")
	return nil
}

func (a *API) ReleaseProducts(req *ReleaseProductRequest, res *ReleaseProductResponse) error {
	var reserved []int
	for _, productCode := range req.ProductCodes {
		err := a.services.WareHouse.ReleaseReservedProducts(req.WarehouseID, productCode)
		if err != nil {
			continue
		}
		reserved = append(reserved, productCode)
	}
	*res = ReleaseProductResponse{
		ReleasedProductCodes: reserved,
	}
	return nil
}

func (a *API) GetRemainingProductCount(req *GetRemainingProductCountRequest, res *GetRemainingProductCountResponse) error {
	products, err := a.services.WareHouse.GetRemainingProductCount(req.WarehouseID)
	if err != nil {
		return err
	}

	*res = GetRemainingProductCountResponse{
		Products: products,
	}

	return nil
}

func (s *Server) Run() error {
	addy, err := net.ResolveTCPAddr("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		return err
	}

	listener := new(API)
	listener.services = s.services
	rpc.Register(listener)
	log.Println("RPC server started")
	for {
		conn, err := inbound.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}
