
# 创建metric server
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml

sudo useradd -m -d /home/work work

sudo passwd work

sudo usermod -aG sudo work


#!/bin/bash

# 设置变量
USERNAME="worker2"
PASSWORD="123456"
SHELL="/usr/bin/zsh"

# 安装 zsh
echo "Installing zsh..."
sudo apt update
sudo apt install -y zsh

# 创建用户并设置默认 shell 为 zsh
echo "Creating user ${USERNAME} with zsh as default shell..."
sudo useradd -m -s ${SHELL} ${USERNAME}
echo "${USERNAME}:${PASSWORD}" | sudo chpasswd

# 禁用自动更新
echo "Disabling automatic updates..."
sudo apt install -y unattended-upgrades
sudo sed -i 's/^APT::Periodic::Update-Package-Lists::/APT::Periodic::Update-Package-Lists:: "0";/' /etc/apt/apt.conf.d/20auto-upgrades
sudo sed -i 's/^APT::Periodic::Unattended-Upgrade::/APT::Periodic::Unattended-Upgrade:: "0";/' /etc/apt/apt.conf.d/20auto-upgrades

# 输出提示信息
echo "User ${USERNAME} created with password ${PASSWORD}."
echo "User ${USERNAME} configured with zsh as default shell."
echo "Automatic updates have been disabled."


#!/bin/bash

# 设置变量
USERNAME="worker"
PASSWORD="ggh_2025"
SUDOERS_LINE="${USERNAME} ALL=(ALL) NOPASSWD: ALL"


sudo mkdir -p /home/${USERNAME}
sudo chmod 700 /home/${USERNAME}
# 创建用户并设置密码
echo "Creating user ${USERNAME}..."
sudo useradd -s /bin/bash ${USERNAME}
echo "${USERNAME}:${PASSWORD}" | sudo chpasswd
sudo cp -r /etc/skel/. /home/${USERNAME}/
sudo chown -R ${USERNAME}:${USERNAME} /home/${USERNAME}/

# 配置 sudoers 文件
echo "Configuring sudoers file for ${USERNAME}..."
sudo bash -c "echo '${SUDOERS_LINE}' >> /etc/sudoers"

# 输出提示信息
echo "User ${USERNAME} created with password ${PASSWORD}."
echo "User ${USERNAME} configured to use sudo without password."

sudo mkdir -p /home/${USERNAME}/.ssh
echo "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDFggCvsMm03XTw5rSkK15AogfxaQT2J65RDxfT8I0lzGydb5ovvF9xJ+MSUi/JT5K0IRklIqzz2oyISnm9VqOUtj8clRiJJqEQuBo8QarN2qD6CIYR6BBXAhpazCr9NSGmqvNEufdC5aoNMqs5F7YNkUZGdHjdB+5eoZmL1+cvKTbFm0YAXsxsoReZUpU2JAOfZn8SfnMNNCMWtpfI4HeqBQncEM6PPlVQxgsaSZjygbCQCpMYJ26T9wWeUIUWeQRFKtffnCQGiuSzYyagwwUmglLjr5HAjCTeo9qJSO2MdU6SFnp9Q44Edma3Yv7XR8BQz3QO/AbHdEvBUoJbqBLnTxMbQDG0wu1uqFRTj8ZKmlMeFCs6JFiwKKcM3Gg9gRqiqlx+taXYWj1N7NTrSVHoGBpO5hbzOlmfeL+fU3iCoR21sb8rgz/iSWFwe5qNMsRn1abkIo4y5XBLMMmYnvJxXmmJGnbpAqgfafym6biAI+fU+glf7VFr7bdr1rwj9Ok= name_fengzi@sina.com" \
  >> /home/${USERNAME}/.ssh/authorized_keys
sudo chown -R ${USERNAME}:${USERNAME} /home/${USERNAME}/



