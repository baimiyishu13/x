#！/bin/bash
wget https://github.com/prometheus/node_exporter/releases/download/v1.7.0/node_exporter-1.7.0.linux-amd64.tar.gz
tar -zxvf node_exporter-1.7.0.linux-amd64.tar.gz -C /usr/local/bin/ --strip-components=1
wget https://github.com/VictoriaMetrics/VictoriaMetrics/releases/download/v1.99.0/vmutils-linux-amd64-v1.99.0.tar.gz

