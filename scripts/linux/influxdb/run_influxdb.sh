docker run -d --name influxdb --restart=always \
-v /home/mqq/influxdb/data:/var/lib/influxdb2 \
-v /home/mqq/influxdb/config.yml:/etc/influxdb2/config.yml \
-p 8086:8086 influxdb:2.0.8