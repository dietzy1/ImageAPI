# syntax=docker/dockerfile:1

FROM golang:1.19

LABEL maintainer="Martin Vad <https://github.com/dietzy1/>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY *.go ./

RUN go build -v -o /imageapi

EXPOSE 8000

CMD [ "/docker-gs-ping" ]
