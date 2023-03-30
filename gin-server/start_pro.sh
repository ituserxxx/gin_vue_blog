#!/bin/bash

ApiCn=blog_api_cn
ImgName=blog_gin_api:v1

ApiCnIsExists=$(docker ps -aqf "name=${ApiCn}")
Img=$(docker images -q --filter reference=${ImgName})
MysqlCn=$(docker ps -qf "name=blog_mysql_cn")

docker-compose -f docker-compose-pro.yml stop

if [ -n "$ApiCnIsExists" ];then
    docker rm -f ${ApiCn}
    echo "$ApiCn ---container remove"
fi

if [ -n "$Img" ];then
  docker rmi -f ${ImgName}
  echo "$ImgName -images remove"
fi
docker build -t ${ImgName} .
docker-compose -f docker-compose-pro.yml up -d
