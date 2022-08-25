# syntax=docker/dockerfile:1

FROM golang:1.19

LABEL maintainer="Martin Vad <https://github.com/dietzy1/>"

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY *.go ./cmd
COPY *.go ./interrnal

RUN go build -o /docker-imageapi

EXPOSE 8000

CMD [ "/docker-imageapi" ]
