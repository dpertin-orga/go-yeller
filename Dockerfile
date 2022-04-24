# syntax=docker/dockerfile:1

FROM golang:alpine

ENV GO111MODULE=on

WORKDIR /go/src/github.com/dpertin-orga/go-yeller/

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY utils ./utils

RUN CGO_ENABLED=0 GOOS=linux go build \
    -a -installsuffix cgo -o /go/bin/goyeller .

FROM alpine:latest

RUN apk add --update ca-certificates # Certificates for SSL
WORKDIR /root/
COPY --from=0 /go/bin/goyeller ./
COPY CHANGELOG.md ./
CMD ["/root/goyeller"]

