FROM golang:latest

RUN go version

copy . /app

Workdir /app


RUN go mod download

RUN go build -o ./bin/main ./cmd/app/app.go

CMD ["/app/bin/main"]
