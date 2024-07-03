# kubectl delete pods 
kubectl delete services publisher rabbitMQ server
kubectl delete deployments publisher rabbitMQ server
kubectl delete secrets tls-certs
kubectl delete configmaps rabbitmq-stable1-conf
kubectl delete -f eks.yaml
kubectl delete -f ./scaling/hpa_server.yaml -f ./scaling/hpa_publisher.yaml
kubectl delete -f ./scaling/vpa_server.yaml -f ./scaling/vpa_publisher.yaml