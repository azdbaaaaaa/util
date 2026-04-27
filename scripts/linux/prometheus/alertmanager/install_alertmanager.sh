#!/bin/bash

set -e
mkdir -p /opt/alertmanager

docker run --name alertmanager -d -p 9093:9093 \
 -v /opt/alertmanager/:/etc/alertmanager/ \
 quay.io/prometheus/alertmanager \
 --config.file=/etc/alertmanager/alertmanager.yml

mkdir -p template
