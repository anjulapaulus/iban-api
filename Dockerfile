# Building
FROM golang:1.16-alpine

WORKDIR /iban_api

COPY go.mod go.sum ./

RUN go mod download

COPY . .