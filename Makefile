# Запуск сервера
run-app:
	docker compose up

migrateup:
	migrate -path migrations -database "postgresql://postgres:password@localhost:5432/lamoda?sslmode=disable" -verbose up


.PHONY: run-app migrateup