kubectl apply -f miniserver.yml
kubectl apply -f miniserver-service.yml
kubectl apply -f ingress.yml
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/mandatory.yaml
