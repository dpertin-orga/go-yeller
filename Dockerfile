FROM golang:latest

WORKDIR /app

ADD . ./

RUN go build -v

CMD ["/app/go-yeller"]
