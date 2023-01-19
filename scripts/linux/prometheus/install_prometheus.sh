#!/bin/bash

docker run  -d --restart=always --name=prometheus\
  -p 9090:9090 \
  -v /opt/prometheus/:/etc/prometheus/  \
  prom/prometheus --web.enable-lifecycle --config.file=/etc/prometheus/prometheus.yml


mkdir -p /opt/grafana/data
chmod 777 -R /opt/grafana/data
docker run -d --restart=always --name=grafana\
  -p 3000:3000 \
  -v /opt/grafana/data:/var/lib/grafana \
  grafana/grafana

# check prometheus config
# docker exec -it prometheus sh -c "promtool check config /etc/prometheus/prometheus.yml"
# curl -X POST http://127.0.0.1:9090/-/reload