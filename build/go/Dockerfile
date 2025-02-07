FROM golang:alpine3.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN apk add build-base

COPY ./ /app/