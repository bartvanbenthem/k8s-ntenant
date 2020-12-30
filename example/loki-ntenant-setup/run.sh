#!/bin/bash

# create namespaces
kubectl create namespace 'co-monitoring'
kubectl create namespace 'team-alpha-dev'
kubectl create namespace 'team-beta-test'

# apply the loki multi tenant setup
kubectl apply -f .

# install Grafana helmchart
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
helm upgrade --install loki --namespace=co-monitoring grafana/loki

# expose grafana to localhost run in seperate terminal
kubectl port-forward svc/grafana -n co-monitoring 3000:3000

# datasource url
http://loki-multi-tenant-proxy.co-monitoring.svc.cluster.local:3100

# team-alpha-dev namespace (pre configured as example)
# beta credentials are generated by credsync service
username: team-alpha-dev
password: alpha


