FROM golang:1.20-alpine
RUN apk add build-base

EXPOSE 8080
WORKDIR /app
COPY ../.. /app

WORKDIR ./cmd/app

RUN go mod download && go mod verify && go build -o .

CMD ["./app"]
