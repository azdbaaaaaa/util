#!/bin/bash

APP="track-service"

# api-versions
kubectl api-versions

# namespace
kubectl create namespace ficool

# configmap
kubectl get configmap -o yaml -n ficool
kubectl delete -n ficool configmap ${APP}

# deployment
kubectl get deployments -n ficool
kubectl create -f deployment.yaml -n ficool
kubectl delete deployment ${APP} -n ficool

kubectl logs -n ficool deployment.apps/${APP}

# service
kubectl get svc -n ficool
kubectl create -f service.yaml -n ficool
kubectl delete svc ${APP} -n ficool

# ingress
kubectl get ing -n ficool
kubectl create -f ingress.yaml -n ficool
kubectl delete ing app-ingress -n ficool
kubectl get ing/app-ingress -n ficool

