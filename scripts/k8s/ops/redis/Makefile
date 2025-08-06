
all: helm install env

helm:
	helm repo add bitnami https://charts.bitnami.com/bitnami
	helm repo update

install:
	#helm install redis bitnami/redis
	#Important: If you install the chart on Minikube, it is not unusual for pods to get stuck in a CrashLoopBackOff event due to unbound volume claims. To prevent this, use the --set option and set persistence.storageClass to nfs-client for both pods.
	helm install redis --set persistence.storageClass=nfs-client,redis.replicas.persistence.storageClass=nfs-client bitnami/redis --set volumePermissions.enabled=true

env:
	export REDIS_PASSWORD=$(kubectl get secret --namespace default redis -o jsonpath="{.data.redis-password}" | base64 --decode)

connect:
	redis-cli -h redis-master.default


	kubectl run --namespace default redis-client --restart='Never'  --env REDIS_PASSWORD=$REDIS_PASSWORD  --image docker.io/bitnami/redis:7.4.0-debian-12-r2 --command -- sleep infinity
	kubectl exec --tty -i redis-client --namespace default -- bash

#2. Connect using the Redis&reg; CLI:
#   REDISCLI_AUTH="$REDIS_PASSWORD" redis-cli -h redis-master
#   REDISCLI_AUTH="$REDIS_PASSWORD" redis-cli -h redis-replicas
#
#To connect to your database from outside the cluster execute the following commands:
#
#    kubectl port-forward --namespace default svc/redis-master 6379:6379 &
#    REDISCLI_AUTH="$REDIS_PASSWORD" redis-cli -h 127.0.0.1 -p 6379