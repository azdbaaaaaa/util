
mkdir -p /opt/grafana/data
mkdir -p /opt/grafana/plugins
chmod 777 -R /opt/grafana
docker run -d --restart=always --name=grafana\
  -p 3000:3000 \
  -v /opt/grafana/data:/var/lib/grafana \
  grafana/grafana