FROM golang:1.23.0-alpine

WORKDIR /src/app

RUN apk add --no-cache git

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod tidy