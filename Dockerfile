# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o bin/server github.com/ecshreve/jepp/cmd/server

EXPOSE 8880

CMD [ "./bin/server" ]