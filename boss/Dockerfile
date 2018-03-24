FROM golang:1.10-stretch

WORKDIR /go/src/github.com/hakobe/grpc-guruguru/boss
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["boss"]