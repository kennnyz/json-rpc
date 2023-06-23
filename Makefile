# Запуск сервера
run-app:
	docker compose up

migrateup:
	migrate -path migrations -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -verbose up


.PHONY: run-app