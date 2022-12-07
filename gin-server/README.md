#gin-server

## 框架搭建
https://github1s.com/fanqingxuan/gin-demo/blob/51830c0f5107c4da921a015cc9403b6e7b9e4fde/middleware/accessLog.go


## redis 库文档

https://github.com/go-redis/redis

## gorm 参考
https://learnku.com/docs/gorm/v2/query/9733#a231f4

## 验证库
https://www.cnblogs.com/jiujuan/p/13823864.html


## 部署Go项目
https://blog.csdn.net/qq_34675369/article/details/112086156




docker run --name nginx-tes1t -p 80:80 -d nginx
docker cp nginx-test:/etc/nginx/conf.d ./docker/nginx/
docker cp nginx-test:/etc/nginx/nginx.conf ./docker/nginx/

docker run -it --name mysql-test -d -p 3309:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:8
docker cp mysql-test:/etc/mysql/conf.d ./docker/mysql/conf
docker cp mysql-test:/var/lib/mysql ./docker/mysql/data


cd 项目dockerfile目录
docker build -t gin_api:v1 .


重启：
docker build -t gin_api:v1 .  && docker-compose up



debug step


docker-compose up 
go run main.go