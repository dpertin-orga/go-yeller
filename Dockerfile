# syntax=docker/dockerfile:1

FROM golang:alpine

ARG VERSION=0.0.1

ENV GO111MODULE=on

WORKDIR /go/src/github.com/dpertin-orga/go-yeller/

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-X main.version=${VERSION}" \
    -a -installsuffix cgo -o /go/bin/goyeller .

FROM alpine:latest

RUN apk add --update ca-certificates # Certificates for SSL
WORKDIR /root/
COPY --from=0 /go/bin/goyeller ./
CMD ["/root/goyeller"]

