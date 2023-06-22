package models

type Product struct {
	Name     string // название товара
	Size     string // размер товара
	Code     string // уникальный код товара
	Quantity int    // количество товара
}

//docker run --name my-postgres-db -e POSTGRES_DB=notifications -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres
