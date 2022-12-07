FROM  golang:latest  as builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"
WORKDIR  /go/src
COPY  . go.mod go.sum config.toml ./
RUN go mod download && go mod tidy &&  go build -o blog_gin_api -ldflags="-w -s" main.go

FROM  alpine:latest
WORKDIR  /root
COPY  --from=builder /go/src/config.toml  .
COPY  --from=builder /go/src/blog_gin_api  .
EXPOSE 6008
CMD  ["./blog_gin_api"]



