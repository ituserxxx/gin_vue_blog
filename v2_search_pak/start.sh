#!/bin/bash
mkdir -p /docker/blog/mysql
cp -r mysql /docker/blog/
docker compose -f docker-compose-env.yml up -d
nohup ./blog_api_v2  &