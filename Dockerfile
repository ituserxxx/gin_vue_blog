FROM golang:alpine  as builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.io,direct"
WORKDIR /app
COPY config.toml go.mod main.go ./
COPY application /app/application
COPY constant /app/constant
COPY library /app/library
COPY middleware /app/middleware
COPY router /app/router
COPY utils /app/utils

RUN  go mod tidy && go build -o blog_gin_api -ldflags "-s -w"  main.go && chmod -R 777 ./blog_gin_api

FROM  alpine:latest
WORKDIR  /root
COPY  --from=builder /app/config.toml  .
COPY  --from=builder /app/blog_gin_api  .
EXPOSE 6008
CMD ["sh", "-c", "/root/blog_gin_api"]