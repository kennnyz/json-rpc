package postgres

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"log"
)

type Warehouse interface {
	ReserveProducts(wareHouseID int, productCodes int) error
	ReleaseReservedProducts(productCodes int) error
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

func (w *WarehouseRepo) ReserveProducts(wareHouseID int, productCode int) error {
	log.Println("Reserving products:", productCode)

	_, err := w.db.Exec("SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;")
	if err != nil {
		return err
	}

	tx, err := w.db.Begin()
	if err != nil {
		return err
	}

	stmtSelect, err := tx.Prepare("SELECT * FROM product_warehouse WHERE warehouse_id = $1 FOR UPDATE;")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmtSelect.Close()

	stmtUpdate, err := tx.Prepare("UPDATE product_warehouse SET reserved = reserved + 1 WHERE warehouse_id = $1 AND product_code = $2;")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmtUpdate.Close()

	rows, err := stmtSelect.Query(wareHouseID)
	if err != nil {
		return err
	}
	rows.Close()

	_, err = stmtUpdate.Exec(wareHouseID, productCode)
	if err != nil {

		if pqErr, ok := err.(*pq.Error); ok {
			errorCode := pqErr.Code
			fmt.Println("Код ошибки:", errorCode)
		} else {
			// Обработка других типов ошибок базы данных
			fmt.Println("Ошибка базы данных:", err)
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (w *WarehouseRepo) ReleaseReservedProducts(productCodes int) error {
	// Реализация освобождения резерва товаров
	log.Println("Releasing reserved products:", productCodes)
	return nil
}

func (w *WarehouseRepo) GetRemainingProductCount(warehouseID int) (int, error) {
	// Реализация получения количества оставшихся товаров на складе
	// Используйте w.DB для выполнения соответствующих операций с базой данных
	log.Println("Getting remaining product count for warehouse:", warehouseID)
	return 0, nil
}
