🎉 接入腾讯云监控，vmagent 会将监控数据 pull 到腾讯云 victoriametrics TSDB

----

#### 前提

1. 下载包

```
wget https://github.com/VictoriaMetrics/VictoriaMetrics/releases/download/v1.99.0/vmutils-linux-amd64-v1.99.0.tar.gz
```

2. 证书文件放置在 `/etc/ssl/certs`

```
cert.pem  key.pem
```

证书：

```sh
cat <<EOF >/etc/ssl/certs/cert.pem
-----BEGIN CERTIFICATE-----

-----END CERTIFICATE-----
EOF

cat <<EOF >/etc/ssl/certs/key.pem
-----BEGIN PRIVATE KEY-----

-----END PRIVATE KEY-----
EOF
```

#### 配置

⚠️注意：`datacenter` 标签用于区分环境 【需添加两个标签】

`/data/vmagent.yml` 

```sh
global:
  external_labels:
    datacenter: heihutao-tw-test # 环境标签1
    heihutao: heihutao-tw-test	 # 环境标签2
scrape_configs:
  - job_name: node-exporter # 监控节点
    static_configs:
    - targets: ['1.1.1.1:9101'，'1.1.1.2:9101']  # 监控节IP
```

`vmagent.service`

```sh
cat <<EOF > /etc/systemd/system/vmagent.service 

[Unit]
Description=vmagent
After=network.target

[Service]
Type=simple
ExecStart=/data/vmagent-prod -tls -tlsCertFile=/etc/ssl/certs/cert.pem -tlsKeyFile=/etc/ssl/certs/key.pem -remoteWrite.bearerToken=** -promscrape.config=/data/vmagent.yml -remoteWrite.url=https://1.15.176.240:8427/api/v1/write
Restart=always

[Install]
WantedBy=multi-user.target

EOF
```

重启服务：

```sh
systemctl daemon-reload
systemctl restart vmagent.service 
systemctl status vmagent.service 
systemctl enable vmagent.service 
```

验证：

Grafana 探索指标 node_uname_info{datacenter=}

