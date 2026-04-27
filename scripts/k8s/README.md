


# 安装
## 准备
1. 安装`kubectl`、`eksctl`、`aws` 等命令
   
## 创建集群
1. 目前都是通过控制台创建的  
    `# eksctl create cluster --name ficool-k8s --region eu-west-1`


2. 更新配置到本地，本地电脑可以连通集群  
    `aws eks update-kubeconfig --region eu-west-1 --name ficool-k8s`
   

3. 创建metris-server
    `kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml`

4. 创建ops命名空间、efs磁盘挂载

```
kubectl create namespace kube-ops
kubectl apply -f https://raw.githubusercontent.com/azdbaaaaaa/util/master/scripts/k8s/ops/storage/storageclass.yaml
kubectl apply -f https://raw.githubusercontent.com/azdbaaaaaa/util/master/scripts/k8s/ops/pv.yaml
kubectl apply -f https://raw.githubusercontent.com/azdbaaaaaa/util/master/scripts/k8s/ops/pvc.yaml
```

5. 创建istio
    `curl -L https://raw.githubusercontent.com/azdbaaaaaa/util/master/scripts/k8s/ops/istio/install.sh | sh -`
   

6. 创建prometheus
   `curl -L https://raw.githubusercontent.com/azdbaaaaaa/util/master/scripts/k8s/ops/prometheus/install.sh | sh -`
