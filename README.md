# micro-todoList
# Go-Micro V4 + RabbitMQ 构造简单备忘录

基于microv4，服务发现使用etcd，支持熔断机制，token验证，网关和各模块之间的rpc通信等

# 项目主要依赖：

**Golang V1.20**

- Gin
- Gorm
- mysql
- go-micro
- protobuf
- grpc
- amqp
- ini
- hystrix
- jwt-go
- crypto

# 项目结构
## 1.micro_todolist 项目总体
```
micro-todolist/
├── app                   // 各个微服务
│   ├── gateway           // 网关
│   ├── task              // 任务模块微服务
│   └── user              // 用户模块微服务
├── bin                   // 编译后的二进制文件模块
├── config                // 配置文件
├── consts                // 定义的常量
├── doc                   // 接口文档
├── idl                   // protoc文件
│   └── pb                // 放置生成的pb文件
├── logs                  // 放置打印日志模块
├── pkg                   // 各种包
│   ├── ctl               // 用户操作
│   ├── e                 // 统一错误状态码
│   ├── logger            // 日志
│   └── util              // 各种工具、JWT等等..
└── types                 // 定义各种结构体
```

## 2.gateway 网关部分
```
gateway/
├── cmd                   // 启动入口
├── http                  // HTTP请求头
├── handler               // 视图层
├── logs                  // 放置打印日志模块
├── middleware            // 中间件
├── router                // http 路由模块
├── rpc                   // rpc 调用
└── wrappers              // 熔断
```

## 3.user && task 用户与任务模块
```
task/
├── cmd                   // 启动入口
├── service               // 业务服务
├── repository            // 持久层
│    ├── db               // 视图层
│    │    ├── dao         // 对数据库进行操作
│    │    └── model       // 定义数据库的模型
│    └── mq               // 放置 mq
├── script                // 监听 mq 的脚本
└── service               // 服务
```


`config/config.ini`文件，直接将 `config.ini.example-->config.ini` 就可以了
conf/config.ini 文件
```ini
[service]
AppMode = debug
HttpPort = :4000

[mysql]
Db = mysql
DbHost = 127.0.0.1
DbPort = 3306
DbUser = micro_todolist
DbPassWord = micro_todolist
DbName = micro_todolist
Charset = utf8mb4

[rabbitmq]
RabbitMQ = amqp
RabbitMQUser = guest
RabbitMQPassWord = guest
RabbitMQHost = localhost
RabbitMQPort = 5672

[etcd]
EtcdHost = localhost
EtcdPort = 2379

[server]
UserServiceAddress = 127.0.0.1:8082
TaskServiceAddress = 127.0.0.1:8083

[redis]
RedisHost = localhost
RedisPort = 6379
RedisPassword = micro_todolist
```


# 运行简要说明
## MacOS
```shell
# 1. 启动环境
make env-up
# 2. 运行服务
make run
```
## Win + Linux
  
```shell
# win
mysql
# linux docker
redis rabbitMQ 
```
  
```shell
# etcd
docker run -d \
  --name etcd \
  --restart on-failure \
  --privileged \
  -p 2379:2379 \
  -e ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379 \
  -e ETCD_ADVERTISE_CLIENT_URLS=http://0.0.0.0:2379 \
  -e ALLOW_NONE_AUTHENTICATION=yes \
  -e ETCD_API=3 \
  -e ETCD_CERT_FILE="/path/to/server.crt" \
  -e ETCD_KEY_FILE="/path/to/server.key" \
  -e ETCD_AUTH="simple" \
  -e ETCD_AUTH_USER="todolist" \
  -e ETCD_AUTH_PASSWORD="todolist" \
  quay.io/coreos/etcd:v3.5.5
```
  
```shell
# etcd-keeper
cd /opt/micro-todoList/etcdkeeper-v0.7.8
./etcdkeeper -h 127.0.0.1 -p 8080
```


