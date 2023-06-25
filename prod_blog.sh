#!/bin/bash
source .prod.env
function create_netwrok(){
  networkIsExists=$(docker network ls --filter name=$network_name -q)
  echo $networkIsExists
  if [ -z "$networkIsExists" ];then
      docker network create $network_name
      echo "$network_name ---network [created succ]"
  fi
}
function run_api_doc_cn(){
  isExists=$(docker ps -aqf "name=$cn_api_doc")
  if [ -n "$isExists" ];then
      docker rm -f $cn_api_doc
      echo "$cn_api_doc---container remove"
  fi
  mkdir -p  $dev_path/api_doc/html
  chmod  -R 755 $dev_path/api_doc
  docker pull registry.cn-shenzhen.aliyuncs.com/star7th/showdoc
  docker tag registry.cn-shenzhen.aliyuncs.com/star7th/showdoc:latest star7th/showdoc:latest
  docker run  -itd \
  --name=$cn_api_doc \
  -p $port_api_doc:80 \
  -v $dev_path/api_doc/html:/var/www/html/ \
  --network=$network_name \
  --restart=on-failure  \
  --user=root \
  --privileged=true \
  star7th/showdoc
  echo "$cn_api_doc ---container [created succ]"
}
function run_mysql_cn() {
  isExists=$(docker ps -aqf "name=$cn_mysql")
  if [ -n "$isExists" ];then
      docker rm -f $cn_mysql
      echo "$cn_mysql ---container [remove]"
  fi
  mkdir -p $dev_path/mysql/data
  mkdir -p $dev_path/mysql/init
  cp $deploy_path/init.sql $dev_path/mysql/init/
  chmod -R 755 $dev_path/mysql

  docker run -itd \
  --name=$cn_mysql \
  -p $port_mysql:3306 \
  --network=$network_name \
  --restart=on-failure  \
  -v $dev_path/mysql/data:/var/lib/mysql \
  -v $dev_path/mysql/my.cnf:/etc/my.cnf \
  -v $dev_path/mysql/init:/docker-entrypoint-initdb.d \
  -e TZ=Asia/Shanghai \
  -e LANG=C.UTF-8 \
  -e MYSQL_DATABASE=$db \
  -e MYSQL_USER=$db_user \
  -e MYSQL_PASSWORD=$db_pass \
  -e MYSQL_ROOT_PASSWORD=$db_root_pass \
  mysql:8.0 \
  --default-authentication-plugin=mysql_native_password \
  --character-set-server=utf8mb4 \
  --collation-server=utf8mb4_general_ci
  echo "$cn_mysql ---container [created succ]"
}
function run_mysql_admin_cn(){
   isExists=$(docker ps -aqf "name=$cn_mysql_admin")
  if [ -n "$isExists" ];then
      docker rm -f $cn_mysql_admin
      echo "$cn_mysql_admin ---container [remove]"
  fi
  docker run -itd \
  --name=$cn_mysql_admin \
  -p $port_mysql_admin:80 \
  --network=$network_name \
  --restart=on-failure  \
  -e PMA_ARBITRARY=0 \
  -e MYSQL_USER=$db_user \
  -e MYSQL_ROOT_PASSWORD=$db_root_pass \
  -e PMA_HOST=$cn_mysql \
  phpmyadmin/phpmyadmin
  echo "$cn_mysql_admin ---container [created succ]"
}
function run_nginx() {
  isExists=$(docker ps -aqf "name=$cn_nginx")
  if [ -n "$isExists" ];then
      docker rm -f $cn_nginx
      echo "$cn_nginx ---container [remove]"
  fi
  chmod -R 755 $dev_path/nginx
  docker run -itd \
  --name=$cn_nginx \
  -p $port_nginx_3:6003 \
  -p $port_nginx_4:6004 \
  --network=$network_name \
  --restart=on-failure  \
  -v $PWD/admin/dist:/app/admin \
  -v $PWD/vue-front/dist:/app/front \
  -v $dev_path/nginx/nginx.conf:/etc/nginx/nginx.conf \
  -v $dev_path/nginx/conf.d:/etc/nginx/conf.d \
  nginx
}
function re_build_api_image(){
  cn_is_exists=$(docker ps -aqf "name=$cn_blog_go_api")
  img_is_exists=$(docker images -q --filter reference=$img_go_api)

  if [ -n "$cn_is_exists" ];then
      docker rm -f $cn_blog_go_api
      echo "$cn_blog_go_api ---container remove"
  fi

  if [ -n "$img_is_exists" ];then
    docker rmi -f $img_go_api
    echo "$img_go_api ----images remove"
  fi
  docker build -f $PWD/gin-server/Dockerfile --no-cache --force-rm --network $network_name -t  $img_go_api $PWD/gin-server/
}
function run_go_api() {
  re_build_api_image
  docker run -itd \
  --name=$cn_blog_go_api \
  -p $port_go_api:6008 \
  --network=$network_name \
  --restart=on-failure  \
  -v /etc/timezone:/etc/timezone \
  -v /etc/localtime:/etc/localtime \
  -v /usr/share/zoneinfo/Asia/:/usr/share/zoneinfo/Asia/ \
  $img_go_api
}
function run_docker_compose(){
  del_all_container
  re_build_api_image
  docker compose --env-file $PWD/.prod.env -f $PWD/docker-compose-prod.yml up -d
}
function stop_all_container() {
  docker stop $cn_api_doc $cn_mysql $cn_mysql_admin $cn_nginx $cn_blog_go_api
  echo "---container [stop all succ]---"
}
function del_all_container() {
  stop_all_container
  docker rm $cn_api_doc $cn_mysql $cn_mysql_admin $cn_nginx $cn_blog_go_api
  docker rmi -f $img_go_api
  echo "---container and image  [rm all succ]---"
}
function run_cn() {
  run_api_doc_cn
  run_mysql_cn
  run_mysql_admin_cn
  run_nginx
  run_go_api
}
function help() {
    echo "params is required, use some ....
--- ps : show all cmd
--- compose : start by docker compose
--- stop : stop all container
--- rm : del all container and images
--- start : start all services
--- restart : [rm] + [start]
    "
}


if [[ "$1" == "stop" ]]; then
 stop_all_container
fi
if [[  "$1" == "rm" ]]; then
  del_all_container
  docker network rm $network_name
fi
if [[  "$1" == "restart" ]]; then
  create_netwrok
  del_all_container
  run_cn
fi
if [[  "$1" == "start" ]]; then
  create_netwrok
  run_cn
fi
if [[  "$1" == "compose" ]]; then
  create_netwrok
  run_docker_compose
fi
if [[  "$1" == "ps" ]]; then
  docker ps --filter network=$network
fi

if [[  "$1" == "" ]];then
  help
fi

