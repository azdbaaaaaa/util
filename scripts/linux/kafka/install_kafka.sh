#!/bin/bash


docker run --restart=always -d --name zookeeper -p 2181:2181 --volume /etc/localtime:/etc/localtime wurstmeister/zookeeper


docker run -d --restart=always --name kafka -p 9092:9092 --link zookeeper \
--env KAFKA_ZOOKEEPER_CONNECT=10.154.0.2:2181 \
--env KAFKA_ADVERTISED_HOST_NAME=10.154.0.2 \
--env KAFKA_ADVERTISED_PORT=9092  \
--volume /etc/localtime:/etc/localtime \
wurstmeister/kafka:2.11-0.11.0.3