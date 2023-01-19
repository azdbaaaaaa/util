
cd /etc/yum.repos.d 
vim mongodb-org-4.0.repo 
# [mongodb-org-4.0]
# name=MongoDB Repository
# baseurl=https://repo.mongodb.org/yum/amazon/2/mongodb-org/4.0/x86_64/
# gpgcheck=0
# enabled=1
# gpgkey=https://www.mongodb.org/static/pgp/server-4.0.asc

yum -y install mongodb-org



# -e MONGO_INITDB_ROOT_USERNAME=root \
# -e MONGO_INITDB_ROOT_PASSWORD=123456 \
# --auth \
docker run -d -p 27017:27017  \
--name mongo \
--restart=always \
-v /opt/mongo/data/db:/data/db \
--cpus="1.0" \
--memory="1g" \
mongo:5.0.10 --tlsAllowInvalidCertificates


docker exec -it mongo mongo --host="mongodb://root:123456@127.0.0.1:27017"

mongo -p 27017



mongodump -h dbhost -d dbname -o dbdirectory
# -h：
# MongoDB 所在服务器地址，例如：127.0.0.1，当然也可以指定端口号：127.0.0.1:27017
# -d：
# 需要备份的数据库实例，例如：test
# -o：
# 备份的数据存放位置，例如：c:\data\dump，当然该目录需要提前建立，在备份完成后，系统自动在dump目录下建立一个test目录，这个目录里面存放该数据库实例的备份数据。
mongorestore -h 127.0.0.1:27017 -d test <path>
# 在dump目录的上级目录执行
# 如果将文档恢复到现有数据库，并且集合和现有文档与要恢复的文档具有相同的 _id 字段值，则 mongorestore 不会覆盖这些文档。

