#!/bin/bash

wget https://github.com/oliver006/redis_exporter/releases/download/v0.30.0/redis_exporter-v0.30.0.linux-amd64.tar.gz
tar xf redis_exporter-v0.30.0.linux-amd64.tar.gz
nohup ./redis_exporter -redis.addr tars-001.kfsfj6.0001.euw1.cache.amazonaws.com:6379 &
nohup ./redis_exporter -redis.addr boss-001.kfsfj6.0001.euw1.cache.amazonaws.com:6379 -web.listen-address :9122 &
nohup ./redis_exporter -redis.addr usersys-001.usersys.kfsfj6.euw1.cache.amazonaws.com:6379 -web.listen-address :9123 &



