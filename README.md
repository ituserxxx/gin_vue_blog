# gin_vue_blog
## 使用 Gin  vue  element 单纯的一个博客
## 前后端分离部署
- 博客前台
- 博客后台管理
- api端

# 部署
首先拉代码
```
git clone https://gitee.com/ituserxxx/gin_vue_blog.git

```

###  docker/docker-compose 一键部署
```
cd  gin_vue_blog/gin-server && start_pro.sh
# 开发模式：cd  gin_vue_blog/gin-server && start_dev.sh
``


```
###  方式 一：前后端独立部署
所需环境
- golang
- mysql
- nginx
- node
- 所需端口：6001  6002  6003 6008


打包
```
# 前台代码
cd vue-front  &&   npm i  && npm  run build  

# 后台代码
cd  admin   && npm  i  && npm  run build:prod

# 编译 api 服务代码
cd  gin-server 
set GOARCH=amd64
set GOOS=linux
go build -o blog_gin_api main.go

# 启动 api 服务
./blog_api_server

```
nginx 配置

```
upstream api {
    server blog_api_cn:6008;
}
server {
    listen 6008;
    location / {
      proxy_pass http://api;
    }
}
server {
    listen 6003;

    location / {
        root   /app/front/;
        index  index.html index.htm;
        try_files $uri $uri/ /index.html;
    }

    location ^~ /blog {
        add_header "Access-Control-Allow-Origin" *;
        try_files $uri $uri/ /index.html;
        proxy_next_upstream http_502 http_504 error timeout invalid_header;
        proxy_set_header Host  $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_pass http://api;
    }

}
server {
    listen 6004;
    location / {
        root   /app/admin/;
        index index.html;
    }

    location ^~ /admin {
        add_header "Access-Control-Allow-Origin" *;
        try_files $uri $uri/ /index.html;
        proxy_next_upstream http_502 http_504 error timeout invalid_header;
        proxy_set_header Host  $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_pass http://api;
    }
}

```

访问地址：
博客前端：ip:6003
博客后台：ip:6004
数据库管理页面：ip:6002


