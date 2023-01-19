#!/bin/bash

# install docker on centos7
yum update -y
yum remove docker  docker-common docker-selinux docker-engine
yum install -y yum-utils device-mapper-persistent-data lvm2 wget
yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
yum install -y docker-ce


# install docker on AWS ec2
yum install -y docker


# start docker
systemctl start docker
systemctl enable docker
docker version

# add user
adduser mqq
# passwd mqq ## tink-news-worker
chmod -v u+w /etc/sudoers
echo "mqq   ALL=(ALL)    NOPASSWD: ALL" >> /etc/sudoers
chmod -v u-w /etc/sudoers


# add user to group
groupadd docker #添加docker用户组 ，如果安装了docker，默认会存在，只需要执行下面的即可
gpasswd -a mqq docker
newgrp docker     #更新用户组


# test mqq 
su - mqq
docker ps -a


# sudo curl -L "https://github.com/docker/compose/releases/download/1.23.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
# sudo chmod +x /usr/local/bin/docker-compose


