FROM golang:alpine  as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

WORKDIR /app
COPY . .
RUN  go mod tidy && go build -o blog_gin_api -ldflags="-w -s" main.go

FROM  alpine:latest
WORKDIR  /root
COPY  --from=builder /app/config.toml  .
COPY  --from=builder /app/blog_gin_api  .
EXPOSE 6008
CMD ["sh", "-c", "/root/blog_gin_api"]