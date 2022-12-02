#编译镜像
FROM  golang:latest  as builder
# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
# 设置 GOPROXY 环境变量
ENV GOPROXY="https://goproxy.cn"
WORKDIR  /go/src
COPY  .  ./
COPY  go.mod  ./
COPY  go.sum  ./
COPY  config.toml  ./

RUN go mod download
RUN go mod tidy
RUN  go build -o blog_gin_api blog_api.go

#产物运行镜像
FROM  alpine:latest
WORKDIR  /root
COPY  --from=builder /go/src/blog_gin_api  .
EXPOSE 6008
CMD  ["./blog_gin_api"]



