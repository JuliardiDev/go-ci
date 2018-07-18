FROM golang:alpine

ADD . /go/src/github.com/juliardidev/go-ci

WORKDIR /go/src/github.com/juliardidev/go-ci

RUN go build main.go

EXPOSE 8080

CMD ["go", "test"]