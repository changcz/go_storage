FROM golang:1.17-alpine AS builder

MAINTAINER changz
WORKDIR /go/src/github.com/flipped-aurora/short-chain/server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy



RUN go build -ldflags="-s -w" -o server .

FROM alpine:latest

LABEL MAINTAINER="changz@heitao.com"

WORKDIR /go/src/github.com/flipped-aurora/short-chain/server

COPY --from=0 /go/src/github.com/flipped-aurora/short-chain/server/server ./

# 暴露9020端口
EXPOSE 9020

CMD ["./server"]