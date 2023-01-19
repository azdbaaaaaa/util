#!/bin/bash

set -e

docker run --name alertmanager -d -p 9093:9093 \
 -v /opt/alertmanager/:/etc/alertmanager/ \
 quay.io/prometheus/alertmanager \
 --config.file=/etc/alertmanager/config.yml


# check prometheus config
# docker exec -it prometheus sh -c "promtool check config /etc/prometheus/prometheus.yml"
# curl -X POST http://127.0.0.1:9090/-/reload