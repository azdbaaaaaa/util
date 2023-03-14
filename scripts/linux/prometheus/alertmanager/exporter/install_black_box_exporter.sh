
docker run --restart=always -d \
  -p 9115:9115 \
  --name blackbox_exporter \
  -v /opt/alertmanager/blackbox/config:/config \
  quay.io/prometheus/blackbox-exporter:latest --config.file=/config/blackbox.yml

echo """
modules:
  http_2xx:
    prober: http
    timeout: 5s
    http:
      method: GET
      preferred_ip_protocol: "ip4"
      tls_config:
        insecure_skip_verify: true
""" > /opt/alertmanager/blackbox/config/blackbox.yml

echo """
- job_name: searchnovels_http2xx_probe
  params:
    module:
    - http_2xx
    target:
    - www.searchnovels.com
  metrics_path: /probe
  static_configs:
  - targets:
    - 172.31.31.238:9115
- job_name: mynovel_http2xx_probe
  params:
    module:
    - http_2xx
    target:
    - www.mynovel.live
  metrics_path: /probe
  static_configs:
  - targets:
    - 172.31.31.238:9115
""" >> /opt/prometheus/prometheus.yml