FROM golang:alpine

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 设置 GOPROXY 环境变量
ENV GOPROXY="https://goproxy.cn"


# 下载全部依赖项
RUN go mod download

# 将我们的代码编译成二进制可执行文件app
RUN go build -o blog_gin_api .

# 移动到用于存放生成的二进制文件的 /dist 目录
WORKDIR /dist

# 将二进制文件从 /build 目录复制到这里
RUN cp /build/blog_gin_api .

# 声明服务端口
EXPOSE 6008

# 启动容器时运行的命令
CMD ["/dist/blog_gin_api"]
