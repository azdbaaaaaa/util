#!/usr/bin/env bash

# metrics see:
# https://awesome-prometheus-alerts.grep.to/rules#kubernetes
# install:
# https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-state-metrics#configuration
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install kube-state-metrics prometheus-community/kube-state-metrics -n kube-ops

#helm uninstall kube-state-metrics


helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
helm install opentelemetry-operator open-telemetry/opentelemetry-operator \
  --set admissionWebhooks.certManager.enabled=false \
  --set admissionWebhooks.certManager.autoGenerateCert.enabled=true \
  --set manager.featureGates=operator.autoinstrumentation.multi-instrumentation \
  -n otel


#helm upgrade --install opentelemetry-operator open-telemetry/opentelemetry-operator \
#--set manager.featureGates=operator.autoinstrumentation.multi-instrumentation