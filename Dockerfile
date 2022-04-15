# syntax=docker/dockerfile:1

FROM golang:alpine
RUN apk add --update ca-certificates # Certificates for SSL
ADD ./go-yeller /go/bin/goyeller
ENTRYPOINT /go/bin/goyeller

