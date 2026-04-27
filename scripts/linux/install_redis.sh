#!/bin/bash 



docker run -d -p 6379:6379 \
--name redis \
--restart=always \
--privileged=true \
-v /opt/redis/conf:/etc/redis/ \
-v /opt/redis/data:/data \
--cpus="1.0" \
--memory="1g" \
redis:6.0.6 redis-server /etc/redis/redis.conf


# wget http://download.redis.io/releases/redis-6.0.6.tar.gz
# bind 127.0.0.1 #注释掉这部分，使redis可以外部访问
# daemonize no#用守护线程的方式启动
# requirepass 你的密码#给redis设置密码
# appendonly yes#redis持久化　　默认是no
# tcp-keepalive 300 #防止出现远程主机强迫关闭了一个现有的连接的错误 默认是300
