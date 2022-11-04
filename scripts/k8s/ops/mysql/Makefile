
all: helm install

helm:
	helm repo add bitnami https://charts.bitnami.com/bitnami
	helm repo update

install:
	helm install mysql bitnami/mysql
