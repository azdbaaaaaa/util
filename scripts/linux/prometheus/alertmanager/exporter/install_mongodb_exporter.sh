#!/bin/bash

docker run --name=mongodb_exporter -d -p 9001:9001 --restart=always \
991619025488.dkr.ecr.eu-west-1.amazonaws.com/mongodb_exporter:latest \
-mongodb.uri=mongodb://ficool:J48Y5WgTplYd@docdb-lighthouse-prd.cluster-ceslrotvh0cw.eu-west-1.docdb.amazonaws.com:27017


docker run --name=mongodb_exporter -d -p 9216:9216 -p 17001:17001 \
percona/mongodb_exporter:0.20 --mongodb.uri=mongodb://172.31.23.144:27017


docker run 991619025488.dkr.ecr.eu-west-1.amazonaws.com/mongodb_exporter:latest -h