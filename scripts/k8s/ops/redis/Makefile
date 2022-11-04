
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