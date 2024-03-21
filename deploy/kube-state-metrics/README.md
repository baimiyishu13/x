👀 适用于当前环境版本，k8s v1.9

前提准备：

1. 准备镜像

```sh
bitnami/kube-state-metrics:1.6.0
```
上传到worker节点,使用sed 明天修改deployment.yaml 中 worker节点名称
```shell
sed -i 's/dev-worker1/worker1/g' deployment.yaml
```

注意事项：`deploy.sh` 文件中

1. Token文件路径：/data/k8s-apiserver_token



测试：IP换成当前主机IP

