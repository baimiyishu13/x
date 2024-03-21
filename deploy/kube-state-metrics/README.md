👀 适用于当前环境版本，k8s v1.9

前提准备：

1. 准备镜像

```sh
bitnami/kube-state-metrics:1.6.0
```
上传到worker1节点,使用sed 修改deployment.yaml 中 worker节点名称
```shell
# kubernetes.io/hostname: tw-worker1
sed -i 's/tw-worker1/worker-hostname/g' deployment.yaml
```

🎉 完成上述步骤再继续

---

部署：

```
kubectl create -f deployment.yaml
```

1.创建serviceaccounts

```sh
kubectl create sa prometheus -n kube-system
```

2.创建prometheus角色并对其绑定cluster-admin

```sh
kubectl create clusterrolebinding prometheus --clusterrole cluster-admin --serviceaccount=kube-system:prometheus
```

3.创建secret; k8s1.24之后默认不会为serveiceaccounts创建secret

```sh
kubectl apply -f - <<EOF
apiVersion: v1
kind: Secret
type: kubernetes.io/service-account-token
metadata:
  name: prometheus-token
  namespace: kube-system
  annotations:
    kubernetes.io/service-account.name: "prometheus"
EOF
```

4.测试访问kube-apiserver

```sh
APISERVER=$(kubectl config view --minify -o jsonpath='{.clusters[0].cluster.server}')
TOKEN=$(kubectl get secret  prometheus-token -n kube-system -o jsonpath='{.data.token}' | base64 --decode)
curl $APISERVER/api --header "Authorization: Bearer $TOKEN" --insecure
```

5.保存token

```sh
echo $TOKEN  > /data/k8s-apiserver_token
```

6.测试访问指标,访问pod性能资源指标：（访问kubelet）

🔔 注意：tw-worker1为当前worker节点的hostname，需要修改

```sh
curl $APISERVER/api/v1/nodes/tw-worker1:10250/proxy/metrics --header "Authorization: Bearer $TOKEN" --insecure
```



配置项：

🔔 sed 修改信息后，复制到监控抓取配置

```sh
sed -i 's/10.61.200.222/master1-IP/g' config.yaml
```

修改为Nodeport端口

```sh
  - job_name: kube-state-metrics
    static_configs:
      - targets: ['1.1.1.1:18566']
```



Grafana ID ：
