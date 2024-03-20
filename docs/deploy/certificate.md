🎉 用于 vmagent传输数据 TLS加密

# 证书
```sh
cat <<EOF > openssl.cnf
[req]
req_extensions = v3_req
distinguished_name = req_distinguished_name
[req_distinguished_name]
[v3_req]
basicConstraints = CA:FALSE
keyUsage = nonRepudiation, digitalSignature, keyEncipherment
subjectAltName = @alt_names
[alt_names]
IP.1 = 1.15.176.240
EOF

openssl req -x509 -nodes -days 3650 -newkey rsa:2048 \
    -keyout key.pem -out cert.pem \
    -subj "/C=CN/ST=shanxi/L=xian/O=IT/CN=monitor" \
    -config openssl.cnf
```
