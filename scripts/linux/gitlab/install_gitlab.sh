#!/bin/bash
set -e

##### 1. 直接安装
# 切换到root
sudo su -


# 添加gitlab用户，添加sudo权限
adduser gitlab
chmod -v u+w /etc/sudoers
echo "gitlab   ALL=(ALL)    NOPASSWD: ALL" >> /etc/sudoers
chmod -v u-w /etc/sudoers

# 切换到gitlab用户
su - gitlab

# install gitlab
wget https://mirrors.tuna.tsinghua.edu.cn/gitlab-ce/yum/el7/gitlab-ce-12.10.6-ce.0.el7.x86_64.rpm --no-check-certificate
sudo yum install -y policycoreutils-python
sudo rpm -i gitlab-ce-12.10.6-ce.0.el7.x86_64.rpm

# 修改配置
sudo vim  /etc/gitlab/gitlab.rb
# 修改配置后，需要配置alb添加http/https监听器以及证书
# external_url 'https://gitlab.ficool.com'
# node_exporter['enable'] = false
# nginx['enable'] = true
# nginx['listen_port'] = 80
# nginx['listen_https'] = false

sudo gitlab-ctl reconfigure
sudo gitlab-ctl restart

# 0 2 * * * /opt/gitlab/bin/gitlab-rake gitlab:backup:create


# cd /var/opt/gitlab/backups
# gitlab-rake gitlab:check
# gitlab-rake gitlab:backup:create
# gitlab-rake gitlab:backup:restore BACKUP=1637373622_2021_11_20_12.10.6


sudo gitlab-rails console -e production
> user=User.where(id:1).first
> user.password='jzOrz8_Vb52SHLdR'
> user.password_confirmation='jzOrz8_Vb52SHLdR'
> user.save!

# root/jzOrz8_Vb52SHLdR


##### 2. 卸载
# cat /opt/gitlab/embedded/service/gitlab-rails/VERSION
# sudo gitlab-ctl stop
# gitlab-ctl cleanse # 保留数据不执行命令
# sudo gitlab-ctl uninstall
# rm -rf /opt/gitlab #保留数据不执行该命令
# sudo yum remove gitlab-ce -y

##### 3. 容器安装
# docker pull gitlab/gitlab-ce
docker run -d --restart=always -p 80:8088 \
--cpus=2 \
-m=8g \
--name gitlab  \
-v /home/gitlab/config:/etc/gitlab \
-v /home/gitlab/logs:/var/log/gitlab \
-v /home/gitlab/data:/var/opt/gitlab \
gitlab/gitlab-ce:12.10.6-ce.0

docker exec -it gitlab update-permissions
docker exec gitlab cp -r /var/opt/gitlab/redis/redis.conf /opt/gitlab/sv/redis/log/config
docker exec gitlab cp -r /var/opt/gitlab/gitaly/config.toml /opt/gitlab/sv/gitaly/log/config
docker exec gitlab cp -r /var/opt/gitlab/postgresql/data/postgresql.conf /opt/gitlab/sv/postgresql/log/config

# docker restart gitlab
# docker logs -f gitlab
# docker stop gitlab && docker rm gitlab


# Stop the processes that are connected to the database
docker exec -it gitlab gitlab-ctl stop puma
docker exec -it gitlab gitlab-ctl stop sidekiq

# Verify that the processes are all down before continuing
docker exec -it gitlab gitlab-ctl status


docker exec -it gitlab gitlab-rake gitlab:check SANITIZE=true
docker exec -it gitlab gitlab-rake gitlab:backup:create
# Run the restore. NOTE: "_gitlab_backup.tar" is omitted from the name
docker exec -it gitlab gitlab-backup restore BACKUP=1637373622_2021_11_20_12.10.6

# Check GitLab

##### Restart the GitLab container and attach config
docker restart gitlab
docker exec gitlab cp -r /var/opt/gitlab/redis/redis.conf /opt/gitlab/sv/redis/log/config
docker exec gitlab cp -r /var/opt/gitlab/gitaly/config.toml /opt/gitlab/sv/gitaly/log/config
docker exec gitlab cp -r /var/opt/gitlab/postgresql/data/postgresql.conf /opt/gitlab/sv/postgresql/log/config
######




