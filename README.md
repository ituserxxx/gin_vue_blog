# gin_vue_blog

体验地址：http://vtian.top/

## 使用 Gin  vue  element 单纯的一个博客
持续更新中
## 前后端分离部署 
- 博客前台  
- 博客后台管理 
- api端



## 部署方式
#### 数据库
- 导入 gin-server 目录下的sql文件入库
- 配置 gin-server/config 中账号密码

#### blog 前台 vue-front

- 进入 vue-front
- 依次执行：npm i 、npm run build
- 将 dist 目录下所有文件上传到服务器你的站点目录下，并配置域名 

nginx 配置如下

```
server{
    listen 80;
    server_name yoursite.com;
    root         /your_blog_dir/;
    location ^~ /blog {
        add_header "Access-Control-Allow-Origin" *;
        proxy_pass http://127.0.0.1:6008;
    }
    location / {
        try_files $uri $uri/ /index.html;
    }
}
```

#### blog 后台  admin

- 进入 admin 目录 
- 依次执行：npm i 、npm run build
- 将 dist 目录下所有文件上传到服务器你的站点目录下，并配置域名 

nginx 配置如下

```
server{
    listen 80;
    server_name admin.yoursite.com;
    root         //your_admin_blog_dir/;
    location ^~ /admin {
        add_header "Access-Control-Allow-Origin" *;
        proxy_next_upstream http_502 http_504 error timeout invalid_header;
        proxy_set_header Host  $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_pass http://127.0.0.1:6008;	
    }
}
```

#### blog api gin-server

- 进入 gin-server目录
- 执行 go mod download
- 打包成linux 可执行文件：

```
     set GOARCH=amd64
     set GOOS=linux
     go build main.go 
```

- 将 生成的 main 二进制文件放到服务器上面，赋予权限：
```
chmod 777 main 
```
- 启动后台执行 
```
nohup ./main &
```

