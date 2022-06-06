#!/bin/bash

docker run -d \
    --name db-simple-blog \
    --restart=always \
    -v $PWD/db:/var/lib/postgresql/data \
    -p 10000:5432 \
    -e POSTGRES_USER=blog \
    -e POSTGRES_PASSWORD=blog \
    -e POSTGRES_DB=blog \
    postgres:alpine