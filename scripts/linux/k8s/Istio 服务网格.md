# Istio 服务网格
see more in `https://istio.io/`

# 为什么使用服务网格

- 为 HTTP、gRPC、WebSocket 和 TCP 流量自动负载均衡。
>  grpc的负载均衡（https://grpc.io/blog/grpc-load-balancing/）

- 通过丰富的路由规则、重试、故障转移和故障注入对流量行为进行细粒度控制。

- 可插拔的策略层和配置 API，支持访问控制、速率限制和配额。

- 集群内（包括集群的入口和出口）所有流量的自动化度量、日志记录和追踪。

- 在具有强大的基于身份验证和授权的集群中实现安全的服务间通信。

## 架构图
![arch-sec](https://uat-docusign-agreement.s3.eu-west-1.amazonaws.com/test/arch-sec.svg)
![arch](https://uat-docusign-agreement.s3.eu-west-1.amazonaws.com/test/arch.svg)
## 1.下载安装
	### 1.1下载
1. [下载页面Istio 1.10.3](!https://github.com/istio/istio/releases/tag/1.10.3)
2.  `cd istio-1.10.3`  
	安装目录包含：  
	- samples/ 目录下的示例应用程序
	- bin/ 目录下的 istioctl 客户端二进制文件
	
3. `export PATH=$PWD/bin:$PATH`


### [1.2安装](https://istio.io/latest/docs/setup/install/istioctl/)
1. istio组件安装   
	`istioctl install -y`
	默认会安装default profile，详见`https://istio.io/latest/docs/setup/additional-setup/config-profiles/`
	安装完成后可以查看安装的组件
	`kubectl -n istio-system get deploy`
	
	默认会安装istiod、ingressgateway、egressgateway。但是一般生产级环境建议只安装istiod，gateway另外单独部署安装。	`istioctl install --set profile=minimal` 

2. 给namespace添加标签，指示istio在部署应用的时候，自动注入envoy边车代理  
  `kubectl label namespace <ficool> istio-injection=enabled`

### 1.3卸载
1. Istio 卸载程序按照层次结构逐级的从 istio-system 命令空间中删除 RBAC 权限和所有资源:  
`istioctl x uninstall --purge` 
2. 命名空间 istio-system 默认情况下并不会被删除。 不需要的时候，使用下面命令删掉它：
`kubectl delete namespace istio-system`
3. 指示 Istio 自动注入 Envoy 边车代理的标签默认也不删除。  
`kubectl label namespace <default> istio-injection-`

### [1.4升级](https://istio.io/latest/docs/setup/upgrade/canary/)

### [1.5多集群安装](https://istio.io/latest/docs/setup/install/multicluster/)

### [更多](https://istio.io/latest/zh/docs/setup/install/)


## 2.架构

## 3.核心概念


### [3.1虚拟服务](https://istio.io/latest/docs/reference/config/networking/virtual-service/)
>定义了一组要在寻址主机时应用的流量路由规则。每个路由规则定义了特定协议流量的匹配标准。如果流量匹配，则将其发送到注册表中定义的命名目标服务（或其子集/版本）。

```
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  namespace: ficool
  name: go-search-service
spec:
  hosts: //流量被发送到的目标主机
    - go-search-service.ficool.svc.cluster.local
  http: //HTTP 流量的路由规则的有序列表。HTTP路由将应用于名为'http- '/'http2- '/'grpc-*'的平台服务端口、使用HTTP/HTTP2/GRPC/ TLS-terminated-HTTPS协议的网关端口和使用HTTP/HTTP2/的服务入口端口GRPC 协议。使用匹配传入请求的第一个规则。
    - name: "uid_0_v3"
      match:
        - headers:
            uid:
              exact: "0"
      route:
        - destination:
            host: go-search-service.ficool.svc.cluster.local
            subset: v3
    - name: "default_v1_v2" // 为调试目的分配给路由的名称。路由的名称将与匹配的名称连接，并将记录在与此路由/匹配匹配的请求的访问日志中。
      route:
        - destination:
            host: go-search-service.ficool.svc.cluster.local
            subset: v2
          weight: 50
        - destination:
            host: go-search-service.ficool.svc.cluster.local
            subset: v1
              weight: 50
      timeout: 60s
      retries:
        attempts: 3
        perTryTimeout: 2s
  exportTo: //此虚拟服务导出到的命名空间列表。导出虚拟服务允许它被其他命名空间中定义的 sidecar 和网关使用。"." 表示导出到相同的命名空间，"*"表示导出到所有命名空间。
    - "."
```


### [3.2目标规则](https://istio.io/latest/docs/reference/config/networking/destination-rule/)
	
> 定义了在路由发生后应用于服务流量的策略。这些规则指定负载平衡的配置、来自 sidecar 的连接池大小以及异常检测设置，以检测和从负载平衡池中驱逐不健康的主机。

```
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  namespace: ficool
  name: go-search-service
spec:
  host: go-search-service.ficool.svc.cluster.local
  trafficPolicy:
    loadBalancer:
      simple: ROUND_ROBIN
  subsets:
    - name: v1
      labels:
        app: go-search-service
        version: v1
    - name: v2
      labels:
        app: go-search-service
        version: v2
      trafficPolicy:
        loadBalancer:
          simple: RANDOM
        connectionPool:
          tcp:
            maxConnections: 100
```

### [3.3网关](https://istio.io/latest/docs/reference/config/networking/gateway/)
>描述了在网格边缘运行的负载均衡器，用于接收传入或传出的 HTTP/TCP 连接。该规范描述了一组应该公开的端口、要使用的协议类型、负载均衡器的 SNI 配置等。



## 4.应用场景

### 4.1配置请求路由
### 4.2故障注入
### 4.3流量转移
### 4.4请求超时
### 4.5重试
### 4.6熔断
### 4.7流量镜像
### 4.8地域负载均衡


<!--#### 动态服务发现
#### 负载均衡
#### HTTP/2 与 gRPC 代理``
#### TLS 终端
#### A/B 测试
#### 金丝雀发布
#### 基于流量百分比切分的概率发布
#### 故障恢复
### 安全-->

## 可观测性
  
### 流量可视化Kiali
> 部署 Kiali 仪表板、 以及 Prometheus、 Grafana、 还有 Jaeger  

~~~
kubectl apply -f samples/addons
kubectl rollout status deployment/kiali -n istio-system
~~~

访问kiali

`istioctl dashboard kiali`

### 




