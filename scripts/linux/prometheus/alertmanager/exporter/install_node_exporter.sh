#!/bin/bash


# wget https://github.com/prometheus/node_exporter/releases/download/v1.1.2/node_exporter-1.1.2.linux-amd64.tar.gz
# tar xvf node_exporter-1.1.2.linux-amd64.tar.gz
# cd node_exporter-1.1.2.linux-amd64
# ./node_exporter &


docker run -d -p 9100:9100 \
  --name node_exporter \
  -v "/proc:/host/proc:ro" \
  -v "/sys:/host/sys:ro" \
  -v "/:/rootfs:ro" \
  --network="host" \
  quay.io/prometheus/node-exporter \
    -collector.procfs /host/proc \
    -collector.sysfs /host/sys \
    -collector.filesystem.ignored-mount-points "^/(sys|proc|dev|host|etc)($|/)"