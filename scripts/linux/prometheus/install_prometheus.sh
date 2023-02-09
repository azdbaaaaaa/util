#!/bin/bash

mkdir -p /opt/prometheus
chmod 777 -R /opt/prometheus
docker run  -d --restart=always --name=prometheus\
  -p 9090:9090 \
  -v /opt/prometheus/:/etc/prometheus/  \
  prom/prometheus --web.enable-lifecycle --config.file=/etc/prometheus/prometheus.yml



# check prometheus config
# docker exec -it prometheus sh -c "promtool check config /etc/prometheus/prometheus.yml"
# curl -X POST http://127.0.0.1:9090/-/reload