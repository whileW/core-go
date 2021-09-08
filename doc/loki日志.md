## 安装
```
wget https://raw.githubusercontent.com/grafana/loki/v2.3.0/cmd/loki/loki-local-config.yaml -O loki-config.yaml
docker run -v /opt/loki/config:/mnt/config -p 3100:3100 --name loki -d grafana/loki:2.3.0 --config.file=/mnt/config/loki-config.yaml
```

## 配置
```
{"log":{"loki":true}}
{"loki":{"addr":"http://127.0.0.1:3100"}}
```