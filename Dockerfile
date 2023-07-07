# syntax=docker/dockerfile:1

FROM golang:1.20 AS builder

WORKDIR /jepp

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd/jepp ./cmd/jepp
COPY ./data/sqlite/jepp.db ./data/sqlite/jepp.db
COPY ./internal ./internal
COPY ./gqlschema ./gqlschema

RUN go build -o bin/jepp github.com/ecshreve/jepp/cmd/jepp

EXPOSE 8082

CMD [ "./bin/jepp" ]

FROM ubuntu:latest
WORKDIR /
COPY --from=builder /jepp/bin/jepp ./
COPY --from=builder /jepp/data/sqlite/jepp.db ./data/sqlite/jepp.db
COPY --from=builder /jepp/internal ./internal

EXPOSE 8082
CMD [ "./jepp" ]