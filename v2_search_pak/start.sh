#!/bin/bash

docker compose -f docker-compose-env.yml up -d
docker compose -f docker-compose-server.yml up -d
go run blog_api_v2