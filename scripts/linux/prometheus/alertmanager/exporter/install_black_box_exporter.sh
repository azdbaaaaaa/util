
docker run --rm \
  -p 9115/tcp \
  --name blackbox_exporter \
  -v /opt/alertmanager/blackbox/config:/config \
  quay.io/prometheus/blackbox-exporter:latest --config.file=/config/blackbox.yml


