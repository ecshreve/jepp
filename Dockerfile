# syntax=docker/dockerfile:1

FROM golang:1.20

WORKDIR /jepp

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd/jepp ./cmd/jepp
COPY ./data/sqlite/jepp.db ./data/sqlite/jepp.db
COPY ./pkg/models ./pkg/models
COPY ./pkg/server ./pkg/server
COPY ./pkg/utils ./pkg/utils
COPY ./docs ./docs
COPY ./static/site ./static/site

RUN go build -o bin/jepp github.com/ecshreve/jepp/cmd/jepp

EXPOSE 8880

CMD [ "./bin/jepp" ]