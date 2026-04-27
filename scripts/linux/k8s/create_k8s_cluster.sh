#创建eks集群

#!/bin/env bash

## https://docs.aws.amazon.com/zh_cn/eks/latest/userguide/install-kubectl.html
## macOS 安装kubectl eksctl
curl -o kubectl https://amazon-eks.s3.us-west-2.amazonaws.com/1.19.6/2021-01-05/bin/darwin/amd64/kubectl
chmod +x ./kubectl
cp ./kubectl /usr/local/bin/kubectl
kubectl version --short --client
## linux
curl -o kubectl https://amazon-eks.s3.us-west-2.amazonaws.com/1.19.6/2021-01-05/bin/linux/amd64/kubectl
chmod +x ./kubectl
cp ./kubectl /usr/local/bin/kubectl
kubectl version --short --client


brew tap weaveworks/tap
brew install weaveworks/tap/eksctl
eksctl version

# 创建集群 目前都是通过控制台创建的
# eksctl create cluster --name ficool-k8s --region eu-west-1 --fargate
# 

# 更新配置到本地，本地电脑可以连通集群
aws eks update-kubeconfig --region eu-west-1 --name prod-ficool
# aws eks update-kubeconfig --region eu-west-1 --name ficool-k8s-01 --role-arn arn:aws:iam::<aws_account_id>:role/<role_name>

# 查看本地k8s环境配置
kubectl config get-contexts

# 创建磁盘挂载 pv && pvc
kubectl apply -f /Users/zhoudongbin/go/src/github.com/azdbaaaaaa/util/scripts/k8s/storage/storageclass.yaml
kubectl apply -f /Users/zhoudongbin/go/src/github.com/azdbaaaaaa/util/scripts/k8s/storage/prod/pv.yaml
kubectl apply -f /Users/zhoudongbin/go/src/github.com/azdbaaaaaa/util/scripts/k8s/storage/prod/pvc.yaml

# 创建metric server
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml


# dashboard token
kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep eks-admin | awk '{print $1}')\n

# ingress not creaated
kubectl logs -n kube-system deployment.apps/aws-load-balancer-controller

# 修改iam用户访问k8s api权限
# https://docs.aws.amazon.com/eks/latest/userguide/add-user-role.html
kubectl describe configmap -n kube-system aws-auth
kubectl edit -n kube-system configmap/aws-auth

# alpine
apk add curl
curl -I -m 10 -o /dev/null -s -w %{http_code} www.baidu.com
# alpine install telnet
apk add busybox-extras
# busybox-extras telnet 127.0.0.1 8991