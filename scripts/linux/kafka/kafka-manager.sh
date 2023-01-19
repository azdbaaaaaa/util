#!/bin/bash

docker run -d --name kafka-manager \
-p 9000:9000 \
--restart=always \
--env ZK_HOSTS=z-1.ficool.jy4av3.c6.kafka.eu-west-1.amazonaws.com:2181,z-3.ficool.jy4av3.c6.kafka.eu-west-1.amazonaws.com:2181,z-2.ficool.jy4av3.c6.kafka.eu-west-1.amazonaws.com:2181 \
sheepkiller/kafka-manager