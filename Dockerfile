FROM golang:latest

ENV GO111MODULE=on
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/

WORKDIR $GOPATH/src/github.com/xuchengyi2015/go-blog
COPY . $GOPATH/src/github.com/xuchengyi2015/go-blog
RUN go build .

EXPOSE 3001
ENTRYPOINT ["./go-blog"]