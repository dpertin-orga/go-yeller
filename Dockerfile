# syntax=docker/dockerfile:1

FROM golang:alpine

ENV GO111MODULE=on

WORKDIR /go/src/github.com/dpertin-orga/go-yeller/

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY utils ./

