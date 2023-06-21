package postgres_repository

import "database/sql"

type Repositories struct {
	Warehouse Warehouse
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Warehouse: NewWareHouseRepo(db),
	}
}
