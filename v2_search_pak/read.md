
# 搜索版本

## 打包后端程序

```shell
cd gin-server
go build -o blog_api_v2 main.go

#给执行权限
chmod +x blog_api_v2
```

## 打包后台
在src/utils/request.js文件中修改为后端接口地址
```shell
cd admin
npm run build:prod
```

## 打包前台
在src/utils/request.js文件中修改为后端接口地址
```shell
cd vue-front
npm run build
```

## 一起放到部署目录下
```shell
cd v2_search_pak
cp -r ../vue-front/dist front
cp -r ../admin/dist admin
```

## 用docker-compose启动mysql数据库/数据库可视化web/美丽搜索服务

```shell
cd v2_search_pak
mkdir -p /docker/blog/mysql
cp -r mysql /docker/blog/
docker compose -f docker-compose-env.yml up -d
```

**或者一键执行 start.sh**
