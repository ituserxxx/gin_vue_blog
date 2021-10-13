#gin-server

## 框架搭建
https://github1s.com/fanqingxuan/gin-demo/blob/51830c0f5107c4da921a015cc9403b6e7b9e4fde/middleware/accessLog.go


## redis 库文档

https://github.com/go-redis/redis

## gorm 参考
https://learnku.com/docs/gorm/v2/query/9733#a231f4

## 验证库
https://www.cnblogs.com/jiujuan/p/13823864.html


## redis容器
docker run -it -d -p 6379:6379  --name=redis --restart=always -v /data/docker-continaer/redis/conf/redis.conf:/etc/redis/redis.conf -v /data/docker-continaer/redis/data:/data  redis redis-server /etc/redis/redis.conf --appendonly yes
## mysql 容器
docker run -it -d -p 3306:3306 --name=mysql  --restart=always -e MYSQL_ROOT_PASSWORD=root -v /data/docker-continaer/mysql/conf:/etc/mysql/conf.d -v  /data/docker-continaer/mysql/logs:/logs -v  /data/docker-continaer/mysql/data:/var/lib/mysql mysql:8 

## 部署Go项目
https://blog.csdn.net/qq_34675369/article/details/112086156


