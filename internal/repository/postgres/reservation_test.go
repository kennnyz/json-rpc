package postgres

import (
	"github.com/kennnyz/lamoda/lamodaTestTask/pkg/database/postgres"
	"log"
	"sync"
	"testing"
)

// потестим резервацию товаров

func TestNewWareHouseRepo(t *testing.T) {
	// подключаемся к базе данных

	// создаем репозиторий

	db, err := postgres.NewClient("host=localhost port=5432 user=postgres password=password dbname=lamoda sslmode=disable timezone=UTC connect_timeout=5")

	if err != nil {
		log.Println(err)
		t.Error(err)
	}

	// создаем репозиторий
	wareHouse := NewWareHouseRepo(db)

	//err = wareHouse.ReserveProducts(2, []int{22, 24, 25})
	//if err != nil {
	//	log.Println(err)
	//	t.Error(err)
	//}

	// запустим несколько горутин которые попытаюстся зарезервировать
	wg := sync.WaitGroup{}

	wg.Add(3)
	for i := 0; i < 3; i++ {

		go func() {
			defer wg.Done()
			nums := []int{22, 24, 25, 22}
			for _, code := range nums {
				err = wareHouse.ReserveProducts(2, code)
				if err != nil {
					log.Println(err)
					t.Error(err)
				}
			}
		}()
	}
	wg.Wait()

}
