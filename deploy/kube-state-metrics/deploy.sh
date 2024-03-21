#!/bin/bash
kubectl get secret  prometheus-token -n default -o jsonpath='{.data.token}' | base64 --decode >> /data/k8s-apiserver_token
TOKEN=$(kubectl get secret  prometheus-token -n default -o jsonpath='{.data.token}' | base64 --decode)
echo $TOKEN
curl https://10.61.200.221:6443/api/v1/nodes/dev-master:10250/proxy/metrics --header "Authorization: Bearer $TOKEN" --insecure