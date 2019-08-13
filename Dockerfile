FROM golang:latest

ENV GO111MODULE=on

WORKDIR $GOPATH/src/github.com/xuchengyi2015/go-blog
COPY . $GOPATH/src/github.com/xuchengyi2015/go-blog
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-blog"]