package postgres_repository

import (
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/kennnyz/lamoda/internal/models"
	"log"
)

type Warehouse interface {
	ReserveProducts(wareHouseID int, productCodes int) error
	ReleaseReservedProducts(warehouseId, productCodes int) error
	GetRemainingProductCount(warehouseID int) ([]models.Product, error)
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

	stmtSelect, err := tx.Prepare("SELECT * FROM product_warehouse WHERE warehouse_id = $1 AND product_code = $2 FOR UPDATE;")
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

	rows, err := stmtSelect.Query(wareHouseID, productCode)
	if err != nil {
		return err
	}

	if !rows.Next() {
		fmt.Println("No products")
		return models.ErrNoProducts
	}

	rows.Close()

	_, err = stmtUpdate.Exec(wareHouseID, productCode)
	if err != nil {

		if pgError, ok := err.(*pgconn.PgError); ok {
			if pgError.Code == "23514" {
				fmt.Println(err)
				return models.ErrNoProducts
			}
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	fmt.Println("success")

	return nil
}

func (w *WarehouseRepo) ReleaseReservedProducts(warehouseID, productCode int) error {
	// Начало транзакции
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}

	stmtSelect, err := tx.Prepare("SELECT * FROM product_warehouse WHERE warehouse_id = $1 AND product_code = $2 FOR UPDATE;")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmtSelect.Close()

	stmtUpdate, err := tx.Prepare("UPDATE product_warehouse SET reserved = reserved - 1, count = count -1 WHERE warehouse_id = $1 AND product_code = $2;")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmtUpdate.Close()

	rows, err := stmtSelect.Query(warehouseID, productCode)
	if err != nil {
		return err
	}

	if !rows.Next() {
		fmt.Println("No products")
		return models.ErrNoProductsInReserve
	}

	rows.Close()

	_, err = stmtUpdate.Exec(warehouseID, productCode)
	if err != nil {
		if pgError, ok := err.(*pgconn.PgError); ok {
			if pgError.Code == "23514" {
				fmt.Println(err)
				return models.ErrNoProducts
			}
		}
		fmt.Println(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	fmt.Println("success")

	return nil
}

func (w *WarehouseRepo) GetRemainingProductCount(warehouseID int) ([]models.Product, error) {
	//получения количества оставшихся товаров на складе
	log.Println("Getting remaining product count for warehouse:", warehouseID)

	var products []models.Product

	query := `SELECT product_name, product_code, count FROM product_warehouse WHERE warehouse_id = $1;`

	rows, err := w.db.Query(query, warehouseID)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.Name, &p.Code, &p.Quantity)
		if err != nil {
			return products, err
		}
		products = append(products, p)
	}

	return products, nil
}
