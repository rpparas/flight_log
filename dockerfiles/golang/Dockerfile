# syntax = docker/dockerfile:1.0-experimental
ARG BUILD_WORKFLOW

FROM golang:1.18-alpine AS base

RUN apk add --no-cache git

WORKDIR /go/src/app
ENV CGO_ENABLED=0

COPY . .

RUN go get -d -v
RUN go install -v .
RUN go mod tidy

FROM base as testing_code
RUN go test -v .

FROM base as production_code
RUN go build -o server main.go

ENTRYPOINT [ "./server" ]

FROM $BUILD_WORKFLOW as final
