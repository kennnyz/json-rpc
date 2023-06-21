package postgres_repository

import (
	"fmt"
	"github.com/kennnyz/lamoda/lamodaTestTask/pkg/database/postgres"
	"log"
	"testing"
)

func TestWarehouseRepo_GetRemainingProductCount(t *testing.T) {
	db, err := postgres.NewClient("host=localhost port=5432 user=postgres password=password dbname=lamoda sslmode=disable timezone=UTC connect_timeout=5")

	if err != nil {
		log.Println(err)
		t.Error(err)
	}

	// создаем репозиторий
	wareHouse := NewWareHouseRepo(db)

	count, err := wareHouse.GetRemainingProductCount(1)
	if err != nil {
		return
	}

	for _, product := range count {
		fmt.Println(product.Name, product.Code, product.Quantity)
	}
}
