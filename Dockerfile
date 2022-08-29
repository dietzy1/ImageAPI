# syntax=docker/dockerfile:1

FROM golang:1.19

LABEL maintainer="Martin Vad <https://github.com/dietzy1/>"

WORKDIR /app
COPY . .

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY *.go /app/

RUN go build -o example cmd/main.go

#EXPOSE 8000

CMD [ "./example" ]
