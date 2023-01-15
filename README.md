# gin_vue_blog
## 使用 Gin  vue  element 单纯的一个博客
## 前后端分离部署 
- 博客前台
- 博客后台管理
- api端

# 部署
首先拉代码
```
git clone xxx.git


###  docker部署

```
cd  gin-server
docker build -t gin_api:v1 .
docker-compose up
``


```
###  方式 一：前后端独立部署
 所需环境
- golang
- mysql
- nginx
- node 

docker run \
-u root \
-itd \
--rm \
--name jenkins \
-p 6005:8080 \
-v jenkins-data:/var/jenkins_home \
-v /var/run/docker.sock:/var/run/docker.sock \
jenkinsci/blueocean


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
go build -o blog_api_server main.go

# 启动 api 服务
./blog_api_server

# 或者后台启动：./blog_api_server  &

```
nginx 配置

```
upstream api {
    server blog_api:6008;
}
server{
    listen 80;
    server_name yoursite.com;

    location ^~ /blog {
        proxy_pass http://api;
    }
    
    location / {
        root   /your_blog_dir/;  #打包后的前台文件放在这个目录下
        index  index.html index.htm;
        try_files $uri $uri/ /index.html;
    }
}
server{
    listen 80;
    server_name admin.yoursite.com;
    root         your_admin_blog_dir/;  #打包后的后台文件放在这个目录下
    location ^~ /admin {
        proxy_next_upstream http_502 http_504 error timeout invalid_header;
        proxy_set_header Host  $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_pass http://api;
    }
}
```

