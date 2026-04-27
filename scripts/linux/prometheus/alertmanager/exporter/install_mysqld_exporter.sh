#!/bin/bash

set -e

docker run -d --name="mysqld-exporter" \
  -p 9104:9104 \
  -e DATA_SOURCE_NAME="admin:tink-news@tcp(database-2.cwmaasjorg3n.eu-west-1.rds.amazonaws.com:3306)/" \
  prom/mysqld-exporter