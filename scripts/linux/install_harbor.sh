#!/bin/bash

yum install -y yum-utils device-mapper-persistent-data lvm2


curl -L https://github.com/docker/compose/releases/download/1.26.2/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose


chmod +x /usr/local/bin/docker-compose

docker-compose version

wget https://github.com/goharbor/harbor/releases/download/v2.6.0/harbor-online-installer-v2.6.0.tgz

tar xvf harbor-online-installer-v2.6.0.tgz