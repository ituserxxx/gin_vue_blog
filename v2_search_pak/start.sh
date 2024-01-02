#!/bin/bash
mkdir -p /docker/blog/mysql
cp -r mysql /docker/blog/
cp ./nginx/vtian.top.conf /etc/nginx/conf.d/
docker compose -f docker-compose-env.yml up -d
kill $(lsof -t -i:6008)
nginx -s reload
nohup ./blog_api_v2  &