package postgres

import (
	"database/sql"
	"log"
)

type Warehouse interface {
	ReserveProducts(wareHouseID int, productCodes []int) ([]int, error)
	ReleaseReservedProducts(productCodes []int) error
	GetRemainingProductCount(warehouseID int) (int, error)
}

type WarehouseRepo struct {
	db *sql.DB // Предполагается, что у вас есть соединение с базой данных
}

func NewWareHouseRepo(db *sql.DB) *WarehouseRepo {
	return &WarehouseRepo{
		db: db,
	}
}

func (w *WarehouseRepo) ReserveProducts(wareHouseID int, productCodes []int) error {
	log.Println("Reserving products:", productCodes)

	// Формируем SQL-запрос для получения и резервирования товаров
	query := `
		UPDATE product
		SET quantity = quantity - 1
		WHERE warehouse_id = ? AND code = ? AND quantity > 0
	`

	// Итерируемся по переданным кодам товаров и резервируем их
	for _, code := range productCodes {
		// Выполняем запрос к базе данных для резервирования товара
		row := w.db.QueryRow(query, wareHouseID, code)

		// Считываем код товара из результата запроса
		var reservedCode int
		err := row.Scan(&reservedCode)
		if err != nil {
			if err == sql.ErrNoRows {
				// Если товар не найден или количество товара равно 0, пропускаем его
				continue
			}
			return err
		}

	}

	// должны добавить в таблицу reservations все резервации

	return nil
}

func (w *WarehouseRepo) ReleaseReservedProducts(productCodes []int) error {
	// Реализация освобождения резерва товаров
	// Используйте w.DB для выполнения соответствующих операций с базой данных
	log.Println("Releasing reserved products:", productCodes)
	return nil
}

func (w *WarehouseRepo) GetRemainingProductCount(warehouseID int) (int, error) {
	// Реализация получения количества оставшихся товаров на складе
	// Используйте w.DB для выполнения соответствующих операций с базой данных
	log.Println("Getting remaining product count for warehouse:", warehouseID)
	return 0, nil
}
