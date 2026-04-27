#!/bin/bash
set -e

curl -L https://packages.gitlab.com/install/repositories/runner/gitlab-ci-multi-runner/script.rpm.sh | sudo bash
sudo yum install -y gitlab-ci-multi-runner


# url and token 可以参照cicd配置里的步骤和token
sudo gitlab-ci-multi-runner register
