#!/bin/bash


du -d1 -h /var/lib/docker/containers | sort -h

sh -c "cat /dev/null > ${log_file}"


# ~/.bash_aliases

# 杀死所有正在运行的容器.
alias dockerkill='docker kill $(docker ps -a -q)'

# 删除所有已经停止的容器.
alias dockercleanc='docker rm $(docker ps -a -q)'

# 删除所有未打标签的镜像.
alias dockercleani='docker rmi $(docker images -q -f dangling=true)'

# 删除所有已经停止的容器和未打标签的镜像.
alias dockerclean='dockercleanc || true && dockercleani'