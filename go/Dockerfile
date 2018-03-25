FROM golang:1.10-stretch

WORKDIR /go/src/github.com/hakobe/grpc-guruguru/go
COPY . .

RUN go get -d -v ./...
RUN go build

CMD ["./go"]