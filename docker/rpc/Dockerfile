FROM golang:1.20-alpine
RUN apk add build-base

EXPOSE 12345
WORKDIR /app
COPY ../.. /app

WORKDIR ./cmd/json-rpc-server

RUN go mod download && go mod verify && go build -o .

CMD ["./json-rpc-server"]
#FROM golang:1.20-alpine
#
#WORKDIR /app
#
#COPY go.mod ./
#COPY go.sum ./
#
#RUN go mod download
#
#COPY . ./
#
#WORKDIR /app
#
#
#
#RUN go build -o ./cmd/json-rpc-server
#
#
#EXPOSE 12345
#
#
#CMD [ "/cmd/json-rpc-server" ]

