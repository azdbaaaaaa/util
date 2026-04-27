#!/usr/bin/env bash

# metrics see:
# https://awesome-prometheus-alerts.grep.to/rules#kubernetes
# install:
# https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-state-metrics#configuration
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

namespace="kube-ops"
# install kube-state-metrics
helm install kube-state-metrics prometheus-community/kube-state-metrics -n ${namespace}

# install prometheus-node-exporter
helm install prometheus-node-exporter prometheus-community/prometheus-node-exporter -n ${namespace}