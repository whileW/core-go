# core-go
> Golang WEB framework 

## 使用
```
go get github.com/whileW/core-go
```

## Components

- [配置](#配置)
  - [默认环境变量配置](#系统内置配置--环境变量)
  - [从file读取配置](#从file读取配置)
  - [从 nacos 配置中心 读取配置](#从nacos配置中心读取配置)
- [grpc](#grpc)
- httpx：req、resp封装，自动保存请求日志
- log：file、loki
- orm：基于配置自动初始化orm实例


### 配置
```
code: conf/*

使用：
获取系统内置配置：conf.GetConf().SysSetting.Env
获取自定义配置: conf.GetStringd("配置名称","默认值")
获取嵌套结构配置：conf.GetChildd("配置名称").GetStringd("配置名称","默认值")
```

### 系统内置配置--环境变量
```
code: conf/env.go

配置项：
ENV: 环境（debug/release）
HTTPADDR: http服务端口
RPCADDR: rpc服务端口
HOST: 本机ip
SYSTEMNAME: 系统名称
CONFFROM: 从哪里读取配置（file\nacos）[默认为file]

使用：
conf.GetConf().SysSetting.Env
```

### 从file读取配置
```
code: conf/file.go

可选配置项（通过环境变量配置）:
CFNAME: 配置文件名称
CFTYPE: 配置文件类型(暂只支持yaml类型的配置文件)
默认为：config.yaml
```

### 从nacos配置中心读取配置
```
code: conf/nacos.go

可选配置项（通过环境变量配置）:
NACOSADDR: nacos的地址--不包括端口（端口默认8848）  
NACOSDATAID: nacos中的dataid(建议程序名称) --如果不配置则采用SYSTEMNAME   
ENV: 环境（nacos中的group）

部署nacos参考: 
github.com/whileW/core-go/conf/nacos-docker-master
https://nacos.io/zh-cn/docs/quick-start-docker.html
https://github.com/nacos-group/nacos-docker
```

### grpc
```
针对小型应用降低grpc使用成本
--利用traefik网关完成类似服务发现、服务注册

配置
traefik：traefik网关地址(127.0.0.1:80)

使用
代码：
import (
    grpc2 "gitee.com/jiransoft/user/pkg/grpc"
    "google.golang.org/grpc"
)

const GrpcName  = "user"  //需要调用的grpc名称
conn, err := grpc.Dial(grpc2.GetTraefikTarget(GrpcName), grpc.WithInsecure(),grpc.WithResolvers(&grpc2.TraefikBuilder{}))

网关配置：
[http]
  [http.routers]
    [http.routers.quotation_grpc]
      service = "product_grpc"
      rule = "Host(`grpc_quotation`)"

  [http.services]
    [http.services.product_grpc]
      [http.services.product_grpc.loadBalancer]
        [[http.services.product_grpc.loadBalancer.servers]]
          url = "h2c://127.0.0.1:30011"
```