#!/bin/bash

app="track-service"
namespace="pre-ficool"

# api-versions
kubectl api-versions

# namespace
kubectl create namespace ${namespace}

# configmap
kubectl get configmap -o yaml -n ${namespace}
kubectl delete -n ${namespace} configmap ${app}
kubectl create configmap ${app} --from-file=${app}.yaml="config/${app}-${env}.yaml" -n ${namespace}

# deployment
kubectl get deployments -n ${namespace}
kubectl create -f deployment.yaml -n ${namespace}
kubectl delete deployment ${app} -n ${namespace}

kubectl logs -n ${namespace} deployment.apps/${app}

# service
kubectl get svc -n ${namespace}
kubectl create -f service.yaml -n ${namespace}
kubectl delete svc ${app} -n ${namespace}

# ingress
kubectl get ing -n ${namespace}
kubectl create -f ingress.yaml -n ${namespace}
kubectl delete ing app-ingress -n ${namespace}
kubectl get ing/app-ingress -n ${namespace}

