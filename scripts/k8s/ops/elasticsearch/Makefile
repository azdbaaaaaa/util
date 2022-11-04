

helm:
	helm repo add elastic https://helm.elastic.co
	helm repo update

install:
	curl -O https://raw.githubusercontent.com/elastic/helm-charts/master/elasticsearch/examples/minikube/values.yaml
	helm install elasticsearch elastic/elasticsearch -f ./values.yaml

	helm install kibana elastic/kibana

	helm install metricbeat elastic/metricbeat