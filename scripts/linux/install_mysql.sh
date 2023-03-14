#!/bin/bash 



sudo docker run -d -p 3306:3306 \
--name mysql \
--privileged=true \
--restart=always \
-v /opt/mysql/conf:/etc/mysql \
-v /opt/mysql/log:/var/log/mysql \
-v /opt/mysql/data:/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=123456 \
--cpus="1.0" \
--memory="1g" \
mysql:5.6

docker update --cpus="2.0" --memory="2g" mysql

docker run -itd --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 --platform linux/x86_64 mysql:5.6

show variables like 'char%';

set character_set_client=utf8;
set character_set_connection=utf8;
set character_set_results=utf8;
set character_set_server=utf8;