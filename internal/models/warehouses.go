package models

type Warehouse struct {
	ID          int             `json:"id"`           // идентификатор склада
	Name        string          `json:"name"`         // название склада
	IsAvailable bool            `json:"is_available"` // признак доступности склада
	Products    map[int]Product `json:"products"`
}
