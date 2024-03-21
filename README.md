# x

🎉 目录说明：

```sh
deploy
├── exporter							# exporter 部署文档
├── kube-state-metrics		# 监控 kubernetes
├── rules									# 告警规则
```

---

参考架构：

![x](./images/README/x.png)

---



监控部署

---

所有均二进制安装：当前支持

 ```
 ➜  x git:(main) ✗ go run main.go install list
 Installable components:
         🎉 vmagent
         🤕 vmalert
         🤕 victoriaMetrics
         🚀 grafnan
         🤕 prometheus
         🤕 alertmanager
         🎉 nodeExporter
         🚀 etcdExporter
         🎉 mysqlExporter
         🤕 mongoExporter
 ```

用法构建：

```
GOOS=linux GOARCH=amd64 go build -o ./bin/
```

准备： 将所需安装二进制文件放在 x 文件的 ./bin 下
示例：

```sh
./x install nodeExporter
```
安装成功显示：
```azure
🎉 Node exporter installed successfully.
```