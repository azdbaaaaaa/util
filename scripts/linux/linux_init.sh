#!/bin/bash

# su - root
sudo su -

yum install -y git

# sudo amazon-linux-extras install -y docker
# install docker on AWS ec2
yum install -y docker

cat > /etc/docker/daemon.json << EOF  
{
    "log-driver":"json-file",
    "log-opts":{
        "max-size":"50m",
        "max-file":"10"
    }
}
EOF

# start docker
systemctl start docker
systemctl enable docker
docker version

## restart docker
systemctl daemon-reload
systemctl restart docker

# add user
adduser worker
# passwd worker ## tink-news-worker
chmod -v u+w /etc/sudoers
echo "worker   ALL=(ALL)    NOPASSWD: ALL" >> /etc/sudoers
chmod -v u-w /etc/sudoers


# add user to group
groupadd docker #添加docker用户组 ，如果安装了docker，默认会存在，只需要执行下面的即可
gpasswd -a worker docker
newgrp docker    #更新用户组

# 
docker run -d --restart=always --name=node_exporter  --net=host --pid=host \
-v "/proc:/host/proc:ro" \
-v "/sys:/host/sys:ro" \
-v "/:/rootfs:ro" \
prom/node-exporter \
--path.procfs=/host/proc \
--path.rootfs=/rootfs \
--path.sysfs=/host/sys \
--collector.filesystem.ignored-mount-points='^/(sys|proc|dev|host|etc)($$|/)'


# install golang 1.16.15
wget https://go.dev/dl/go1.16.15.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.15.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile
source /etc/profile

# install amazon-efs-utils
yum install -y amazon-efs-utils
pip3 install botocore

su - worker

# set aws config
sudo aws configure set aws_access_key_id AKIA6NYIZRJIO4ST4QOR
sudo aws configure set aws_secret_access_key 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
sudo aws configure set default.region eu-west-1

# create a new
#ssh-keygen -t rsa
# create gitlab-runner ssh key
mkdir .ssh
chmod 700 .ssh
ssh-keygen -f ".ssh/id_rsa" -N ""
echo "" >> /home/worker/.ssh/authorized_keys
chmod 600 .ssh/authorized_keys

# mount
sudo mkdir -p /log
sudo mount -t efs fs-6177e455 /log
sudo chown worker:worker /log

# 
aws configure set aws_access_key_id AKIA6NYIZRJIO4ST4QOR
aws configure set aws_secret_access_key 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
aws configure set default.region eu-west-1

