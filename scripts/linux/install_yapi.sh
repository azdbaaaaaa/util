#!/bin/bash
set -e

docker volume create yapi-mongo
docker run -d --restart=always --name yapi-mongo -p 127.0.0.1:27017:27017  -v /log/data/mongodb:/data/db mongo
docker run -d --restart=always --name yapi -p 3100:3000 --link yapi-mongo crper/yapi



{
  "port": "3000",
  "adminAccount": "zhoudongbin@searchnovels.com",
  "timeout":120000,
  "db": {
    "servername": "127.0.0.1",
    "DATABASE": "yapi",
    "port": 27017
  }
}

docker run -it --rm \
  --entrypoint npm \
  --workdir /yapi/vendors \
  -v /opt/yapi/config.json:/yapi/config.json \
  registry.cn-hangzhou.aliyuncs.com/anoyi/yapi \
  run install-server


docker run -d \
  --name yapi \
  --workdir /yapi/vendors \
  -p 3000:3000 \
  -v /opt/yapi/config.json:/yapi/config.json \
  registry.cn-hangzhou.aliyuncs.com/anoyi/yapi \
  server/app.js  