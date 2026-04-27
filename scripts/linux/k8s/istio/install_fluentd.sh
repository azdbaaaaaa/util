#!/bin/bash


# https://www.shebanglabs.io/logging-with-efk-on-aws-eks/
# https://github.com/kiwigrid/helm-charts/tree/master/charts/fluentd-elasticsearch

helm repo add kokuwa https://kokuwaio.github.io/helm-charts

helm install kokuwa/fluentd-elasticsearch --namespace kube-system -f fluentd.yaml --generate-name

helm list -n kube-system

helm uninstall fluentd-elasticsearch-1637834594 -n kube-system

helm install kokuwa/fluentd-elasticsearch --namespace kube-system -f fluentd.yaml --generate-name




# docker run --rm -v ~/.aws:/root/.aws -p 9200:9200 abutaha/aws-es-proxy:v1.0 -endpoint https://vpc-ficool-test-6hzr7wzi3uptwidursdnki2nyy.eu-west-1.es.amazonaws.com -listen 0.0.0.0:9200