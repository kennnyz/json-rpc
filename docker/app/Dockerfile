FROM golang:1.20-alpine
RUN apk add build-base

EXPOSE 8080
WORKDIR /app
COPY ../.. /app

WORKDIR ./cmd/app

RUN go mod download && go mod verify && go build -o .

CMD ["./app"]

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
#RUN go build -o ./cmd/app
#
#
#EXPOSE 8080
#
#CMD [ "/cmd/app" ]

#After write Dockerfile:

#Build docker with tag:
#docker build --tag docker-gs-ping .
#
#Run docker container with publish port:
#docker run --publish 8080:8080 docker-gs-ping