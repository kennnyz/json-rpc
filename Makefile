# Запуск сервера
run-app:
	go run json-rpc-server/main.go

run-server:

build:
	go build -o app ./cmd/app/app.go
	go build -o rpcServer ./cmd/app/app.go



create-database:
	docker run -d --name lamoda -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=lamoda -d postgres

.PHONY: run-app create-database
