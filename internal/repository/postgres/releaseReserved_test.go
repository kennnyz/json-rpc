package postgres

import (
	"github.com/kennnyz/lamoda/lamodaTestTask/pkg/database/postgres"
	"log"
	"testing"
)

func TestWarehouseRepo_ReleaseReservedProducts(t *testing.T) {
	db, err := postgres.NewClient("host=localhost port=5432 user=postgres password=password dbname=lamoda sslmode=disable timezone=UTC connect_timeout=5")

	if err != nil {
		log.Println(err)
		t.Error(err)
	}

	// создаем репозиторий
	wareHouse := NewWareHouseRepo(db)

	err = wareHouse.ReleaseReservedProducts(2, 22)
}
