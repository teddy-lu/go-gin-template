# go-gin-demo

## 项目信息

### 背景

- 为了有一套适合自己开发的GO项目Demo，所以自己动手写了一个项目框架。
- 项目旨在开箱即用，不需要重复造框架
- 项目旨在为其他开发者提供参考，帮助开发者快速开发项目
- 先fork为新项目，然后修改项目名，再修改项目描述，按照自己的项目修改config，最后修改项目地址，提交git，就可以开发了

### 项目结构

```text
├── cmd # 项目入口
│   └── main.go
├── config # 项目配置
│   └── config.go
├── docs # 文档
│   ├── README.md # 项目md文档
│   └── ...
├── go.mod
├── go.sum
├── internal # 项目核心代码
│   ├── dao # 数据访问层
│   │   └── dao.go
│   ├── models # 数据模型层
│   │   └── model.go
│   ├── service # 业务逻辑层
│   │   └── service.go
│   └── api # 接口服务层
│       └── index.go
├── pkg # 项目工具包
│   └── logger
│       └── logger.go # 日志模块
├── config.example.yaml # 配置文件示例
├── .air.toml # 热重载配置文件
├── routers
│   └── route.go # 路由配置
├── utils # 项目函数包
│   └── helpers.go # 辅助函数
```

### 项目依赖

- github.com/fsnotify/fsnotify
- github.com/gin-gonic/gin
- github.com/go-redis/redis
- gorm.io/driver/mysql
- github.com/lestrrat-go/file-rotatelogs
- go.uber.org/zap
- github.com/spf13/pflag
- github.com/spf13/viper
- github.com/cosmtrek/air
- gorm.io/gorm

### 项目启动

```shell
# 启动方式1
go run cmd/main.go -c config.yaml

# 启动方式2（热重载）下面的网址可以查看热重载的原理
# https://github.com/cosmtrek/air/blob/master/README-zh_cn.md?plain=1
air
```

### 项目配置文件

```yaml
name: go-gin-demo
mode: debug
addr: 8080

db: # mysql数据库配置
  name:
  host: 127.0.0.1
  port: 3306
  username:
  password:
  charset: utf8mb4
  max_idle_cons: 10
  max_open_cons: 2
redis: # redis配置
  host: 127.0.0.1
  port: 6379
  auth:
  db: 0

log:
  level: debug # 日志级别，info，debug，error
  file_format: "%Y%m%d" # 文件格式
  max_save_days: 30
  file_type: one #one, level 单文件存储还是以level级别存储
```