# syntax=docker/dockerfile:1

FROM golang:1.20 AS builder

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

FROM ubuntu:latest
WORKDIR /
COPY --from=builder /jepp/bin/jepp ./
COPY --from=builder /jepp/data/sqlite/jepp.db ./data/sqlite/jepp.db
COPY --from=builder /jepp/docs ./docs
COPY --from=builder /jepp/static/site ./static/site
COPY --from=builder /jepp/pkg ./pkg

EXPOSE 8880
CMD [ "./jepp" ]