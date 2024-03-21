🎉 cassandra 监控社区并不活跃

在cassandra 节点 - data1

1. 下载包

```sh
wget https://repo1.maven.org/maven2/io/prometheus/jmx/jmx_prometheus_javaagent/0.19.0/jmx_prometheus_javaagent-0.19.0.jar
```

2. 修改配置：

```sh
cat <<EOF >> /etc/cassandra/conf/cassandra-env.sh
# 7070为暴露的数据采集端口
JVM_OPTS="\$JVM_OPTS -javaagent:\$CASSANDRA_HOME/lib/jamm-0.3.0.jar -javaagent:\$CASSANDRA_HOME/lib/jmx_prometheus_javaagent-0.19.0.jar=7070:/etc/cassandra/conf/cassandra-jmx.yaml"
EOF
```

3. 重启服务：

```sh
systemctl restart cassandra
systemctl status cassandra
```

4. 验证：

```curl
 curl 127.0.0.1:7070/metrics
```



