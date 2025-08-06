#!/bin/bash
set -e

# 配置变量
USERNAME="worker"
PASSWORD="ggh_2025"
SSH_KEY="ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDFggCvsMm03XTw5rSkK15AogfxaQT2J65RDxfT8I0lzGydb5ovvF9xJ+MSUi/JT5K0IRklIqzz2oyISnm9VqOUtj8clRiJJqEQuBo8QarN2qD6CIYR6BBXAhpazCr9NSGmqvNEufdC5aoNMqs5F7YNkUZGdHjdB+5eoZmL1+cvKTbFm0YAXsxsoReZUpU2JAOfZn8SfnMNNCMWtpfI4HeqBQncEM6PPlVQxgsaSZjygbCQCpMYJ26T9wWeUIUWeQRFKtffnCQGiuSzYyagwwUmglLjr5HAjCTeo9qJSO2MdU6SFnp9Q44Edma3Yv7XR8BQz3QO/AbHdEvBUoJbqBLnTxMbQDG0wu1uqFRTj8ZKmlMeFCs6JFiwKKcM3Gg9gRqiqlx+taXYWj1N7NTrSVHoGBpO5hbzOlmfeL+fU3iCoR21sb8rgz/iSWFwe5qNMsRn1abkIo4y5XBLMMmYnvJxXmmJGnbpAqgfafym6biAI+fU+glf7VFr7bdr1rwj9Ok= name_fengzi@sina.com"

# 创建用户及目录
echo "👤 创建用户 ${USERNAME}..."
sudo useradd -s /bin/bash -m -d "/home/${USERNAME}" "${USERNAME}" 2>/dev/null || true
echo "${USERNAME}:${PASSWORD}" | sudo chpasswd
sudo chmod 700 "/home/${USERNAME}"

# 安全配置sudo权限
echo "🛡️ 配置sudo权限..."
if ! sudo grep -q "${USERNAME} ALL=(ALL) NOPASSWD: ALL" /etc/sudoers; then
    echo "${USERNAME} ALL=(ALL) NOPASSWD: ALL" | sudo tee -a /etc/sudoers >/dev/null
fi

# 配置SSH密钥
echo "🔑 设置SSH密钥..."
sudo mkdir -p "/home/${USERNAME}/.ssh"
echo "${SSH_KEY}" | sudo tee "/home/${USERNAME}/.ssh/authorized_keys" >/dev/null
sudo chown -R "${USERNAME}:${USERNAME}" "/home/${USERNAME}/.ssh"
sudo chmod 700 "/home/${USERNAME}/.ssh"
sudo chmod 600 "/home/${USERNAME}/.ssh/authorized_keys"

# 安装基础工具
echo "🧱 安装基础工具..."
sudo yum update -y
sudo yum install -y yum-utils curl wget git zsh

# 安装Docker
echo "🐳 安装Docker..."
sudo yum remove -y docker docker-client docker-client-latest docker-common \
    docker-latest docker-latest-logrotate docker-logrotate docker-engine podman buildah 2>/dev/null || true

sudo yum install -y yum-utils device-mapper-persistent-data lvm2
sudo yum-config-manager --add-repo https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

sudo yum install -y docker-ce docker-ce-cli containerd.io

# 配置Docker
echo "⚙️ 配置Docker..."
sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<EOF
{
  "data-root": "/var/lib/docker",
  "live-restore": true,
  "log-driver": "json-file",
  "log-opts": {"max-file": "3", "max-size": "10m"},
  "registry-mirrors": [
    "https://dockerpull.cn",
    "https://docker.mirrors.ustc.edu.cn",
    "https://hub-mirror.c.163.com",
    "https://mirror.baidubce.com"
  ]
}
EOF

# 启动Docker服务
echo "🚀 启动Docker服务..."
sudo systemctl enable docker
sudo systemctl daemon-reload
sudo systemctl restart docker
sudo usermod -aG docker "${USERNAME}"

# 运行node-exporter
docker run -d   --name node_exporter   --restart unless-stopped   --pid=host   --network=host   -v "/proc:/host/proc:ro"   -v "/sys:/host/sys:ro"   -v "/:/rootfs:ro"   quay.io/prometheus/node-exporter   --path.procfs=/host/proc   --path.sysfs=/host/sys   --path.rootfs=/rootfs   --collector.filesystem.ignored-mount-points="^/(sys|proc|dev|host|etc)($|/)"   --collector.diskstats   --collector.filesystem   --collector.cpu   --collector.meminfo   --collector.loadavg   --collector.processes

# 安装Node.js
echo "⬇️ 安装Node.js 18..."
curl -fsSL https://rpm.nodesource.com/setup_18.x | sudo bash -
sudo yum install -y nodejs

# 安装PM2
echo "📦 安装PM2..."
sudo npm install -g pm2

# 输出验证信息
echo -e "\n✅ 安装完成！"
echo "----------------------------------------"
echo "用户信息:"
echo "用户名: ${USERNAME}"
echo "密码: ${PASSWORD}"
echo "SSH密钥已配置"
echo "----------------------------------------"
echo "软件版本:"
docker --version 2>/dev/null || echo "Docker未安装"
node --version 2>/dev/null || echo "Node.js未安装"
npm --version 2>/dev/null || echo "npm未安装"
pm2 --version 2>/dev/null || echo "PM2未安装"
echo "----------------------------------------"