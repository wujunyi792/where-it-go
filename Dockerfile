FROM golang:1.16-alpine AS builder

RUN go env -w GO111MODULE=auto \
  && go env -w CGO_ENABLED=0 \
  && go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /build

COPY . .

RUN set -ex \
    && cd /build \
    && pwd \
    && ls \
    && go build -ldflags "-s -w -extldflags '-static'" -o App ./cmd/main.go

FROM alpine:latest

WORKDIR /Serve
COPY ./ ./

COPY --from=builder /build/App ./App

RUN  echo 'http://mirrors.ustc.edu.cn/alpine/v3.5/main' > /etc/apk/repositories \
    && echo 'http://mirrors.ustc.edu.cn/alpine/v3.5/community' >>/etc/apk/repositories \
    && apk update && apk add tzdata \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

ENTRYPOINT [ "/Serve/App" ]