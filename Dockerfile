# syntax=docker/dockerfile:1

FROM golang:1.20

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd/server ./cmd/server
COPY ./data/sqlite ./data/sqlite
COPY ./pkg ./pkg
COPY ./docs ./docs
COPY ./static/site ./static/site

RUN go build -o bin/server github.com/ecshreve/jepp/cmd/server

EXPOSE 8880

CMD [ "./bin/server" ]