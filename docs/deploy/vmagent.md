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
MIIDLjCCAhagAwIBAgIJANBvusH+Li+GMA0GCSqGSIb3DQEBCwUAMEwxCzAJBgNV
BAYTAkNOMQ8wDQYDVQQIDAZzaGFueGkxDTALBgNVBAcMBHhpYW4xCzAJBgNVBAoM
AklUMRAwDgYDVQQDDAdtb25pdG9yMB4XDTI0MDMyMDAyNDMxNVoXDTM0MDMxODAy
NDMxNVowTDELMAkGA1UEBhMCQ04xDzANBgNVBAgMBnNoYW54aTENMAsGA1UEBwwE
eGlhbjELMAkGA1UECgwCSVQxEDAOBgNVBAMMB21vbml0b3IwggEiMA0GCSqGSIb3
DQEBAQUAA4IBDwAwggEKAoIBAQCqFhIR0H8ZHVkTQ7HRx6zOuu+twnnRAvx5esp/
mFBly4hF6J4MfI2fzLLf2lSIMOhIRnicBmK65n270TDLQheheF2+1EAlVkGNzJW0
DmZkAEFouYxsuUIP6Bx4TA+R5iMX3Yo3EKh8rIM44OtjtwFtma6nNako0ftKuBU0
CrbW2ftL1FHVtLxEVe0YFuymztIRuEGxkLCBzhps37QesZdCmLKPmHUNux+S+Qd9
/mkfT6ncNCopk33gbzDYwTsxhQ3vmqBpnfH4CFgjnCzxYaa2qsSlGG9XXuEKSMVl
ZlWrP87Q2xa+kfC89QUsv/JVyQsHKmMrf+6d+tzzCokxn1v3AgMBAAGjEzARMA8G
A1UdEQQIMAaHBAEPsPAwDQYJKoZIhvcNAQELBQADggEBAGV5UIjy8+y1jYvCFGEo
c6dRsU+7Y/WIpe8N+zxzrOBMXdXiiJja5+isDQNyQa0tDQPT3qddxGoJnQHPqVne
Wj+fQaTvvnQ0Pvak9LG18Z0lgcmy7F7ZnNanePYxgqe+x5w6USYzmTDtysLWJnV4
ICSQDtv728SL0Iqpwn/xQ5X1nsdGEi/OzQipZTkblCyJRRGPrm0OP+CMDB5HUc1h
IfDGnaXnZqcqZ2TLR8ICLE4bMW6/0FLBLMCTzdhPqppxlMQzlWL9ZYOtqJhjwX6v
q9Qz727ugtdOiXq5K1DlnPhncWkpZiKQixWwse9id1RGGjlr4fZahOh6I2ax3s2+
C60=
-----END CERTIFICATE-----
EOF

cat <<EOF >/etc/ssl/certs/key.pem
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCqFhIR0H8ZHVkT
Q7HRx6zOuu+twnnRAvx5esp/mFBly4hF6J4MfI2fzLLf2lSIMOhIRnicBmK65n27
0TDLQheheF2+1EAlVkGNzJW0DmZkAEFouYxsuUIP6Bx4TA+R5iMX3Yo3EKh8rIM4
4OtjtwFtma6nNako0ftKuBU0CrbW2ftL1FHVtLxEVe0YFuymztIRuEGxkLCBzhps
37QesZdCmLKPmHUNux+S+Qd9/mkfT6ncNCopk33gbzDYwTsxhQ3vmqBpnfH4CFgj
nCzxYaa2qsSlGG9XXuEKSMVlZlWrP87Q2xa+kfC89QUsv/JVyQsHKmMrf+6d+tzz
Cokxn1v3AgMBAAECggEATFGw8/tpPHy1vnYusSH2waCR8ZcFECtV3LCjcpOG59Zn
JG9Xk+TDq8OvM9EFA0Nmxx+P+PrIjHLmUkNwsJPMhLC85+bXjalhqt31AqE/gnm2
3+X6Q8LxCLAvLi97AQ2SC6Bl54V7BM5n5zrNhKXWZzaBbxgd+moWZxWotfxkxIy5
hA1nXhA1/Mr+1leqYNdWs+JGkgS+6SoNoFv/q0AO/gFIL71jCCzBJANGU04v/pdV
1YBhjddoeetmkRfBAhMsWK5RO8Sd3EJKgoNoaQmAdesGvzT0q9BwZdZ2Liv8aVyK
1zP2n8rbWaBrw9KpknmQw3j4wjh60nAcnaRQydMSmQKBgQDfxvNl3ljcZeoVI4Ce
msrmicqBHq2iBTCk8Km2S4vTk6Arz/ns/Z/6/S07yQ/gO2HUNi5YwotKb2d3EwnX
tFvLw9IreL0Qi22IriwcXUaXhW5hy0S6mnVd73VJYltoTJ7rPgXpxP948hW132ts
AJ03UB9tnY75yBkkz7X2rzX6ZQKBgQDCk+wYiKxbBjSD5J31mHZQJ6Wma4/VXxG4
tGzM7tv1MwVSV30l8rN1by8YbCimQXCPXMp7PjroPchR515iNfDM27I5U1aC5K1Y
obAlXPmgUQqM8r6UO+NoZOkOZ77UAfxjx8WjVfVebe03tVI/yqQL8WOGqcvlp2mu
Dy2NCFHJKwKBgCwaFZYaAAFPxJZt7MdVUm+k8FgKJ8Yqp6+aDphywxfrnEGGN4dR
ZNoeU2/Y7FwQ1/LuyquLi9AbrJR3GBV+4iiCa0VOAlTkZa3uRZtBdswd6+cyHFV6
Y28j5kWUMNeZYiiSjK/Jt1+qtkSlx5H9fP6Nt+PujWME4I+0r/PyC/1RAoGBAJgi
IgHYsrIEuwC4ukxxJWhsZ1ckWHRi/a8l2r+srmwbtTh5UU9fYLMb+H2m6216UvC0
g1gzAsncrIhlV2WAeUMjL0klAMoc0/Atk1TTShXv50jm7t1lqdtQsuvhb9HBr72w
T7XymoN3fXcGHUXlQbhYxxFeJx7rc3R1R48yhwe7AoGAQ3F1pvyyOICMqLbWdU9n
z2RRYGVd4umMdglq5kdn01wh3BIfg14OGldNWCHoIaZ6TdjXEy1a+6HkcJTNtEfC
I6HD34hbveaE8AiEEzPCPoKPbLIWjJOpQst5gt2JNTJ00O/YJ+1rp7/zNSloWdsw
059MLAa/GJ/SAl1u9TlQYCI=
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
ExecStart=/data/vmagent-prod -tls -tlsCertFile=/etc/ssl/certs/cert.pem -tlsKeyFile=/etc/ssl/certs/key.pem -remoteWrite.bearerToken=5Q2MhAoVIEBo2wWgsWSAICJbWQPYXmRyQYPLPFdHT8qFG9cLbBRs0EwwY3xqw8gD -promscrape.config=/etc/prometheus/prometheus.yml -remoteWrite.url=https://1.15.176.240:8427/api/v1/write
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

