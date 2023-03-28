#!/bin/bash
source .dev.env
function create_netwrok(){
  networkIsExists=$(docker network ls --filter name=$network_name -q)
  echo $networkIsExists
  if [ -z "$networkIsExists" ];then
      echo "$network_name ---network create"
      docker network create $network_name
  fi
}
function start_api_doc(){
  isExists=$(docker ps -aqf "name=$cn_api_doc")
  if [ -n "$isExists" ];then
      docker rm -f $cn_api_doc
      echo "$cn_api_doc ---container remove"
  fi
  mkdir -p  $data_path/api_doc/html && chmod  -R 777 $data_path/api_doc/
  docker pull registry.cn-shenzhen.aliyuncs.com/star7th/showdoc
  docker tag registry.cn-shenzhen.aliyuncs.com/star7th/showdoc:latest star7th/showdoc:latest
  docker run  -itd \
  --name=$cn_api_doc \
  -p $port_api_doc:80 \
  -v $data_path/api_doc/html:/var/www/html/ \
  --network=$network_name \
  --restart=on-failure  \
  --user=root \
  --privileged=true \
  star7th/showdoc
}
function start_mysql() {
  isExists=$(docker ps -aqf "name=$cn_mysql")
  if [ -n "$isExists" ];then
      docker rm -f $cn_mysql
      echo "$cn_mysql ---container remove"
  fi
  mkdir -p $data_path/mysql/data
  mkdir -p $data_path/mysql/my.cnf
  mkdir -p $data_path/mysql/init
  mkdir -R 777 $data_path/mysql/
  cp $data_path/init.sql $data_path/mysql/init/

  docker run -itd \
  --name=$cn_mysql \
  -p port_mysql:3306 \
  -v $data_path/mysql/data:/var/lib/mysql \
  -v $data_path/mysql/my.cnf:/etc/my.cnf \
  -v $data_path/mysql/init:/docker-entrypoint-initdb.d \
  -e TZ=Asia/Shanghai \
  -e LANG=C.UTF-8 \
  -e MYSQL_DATABASE=$db \
  -e MYSQL_USER=$db_user \
  -e MYSQL_PASSWORD=$db_pass \
  -e MYSQL_ROOT_PASSWORD=$db_root_pass \
}

create_netwrok
start_api_doc

