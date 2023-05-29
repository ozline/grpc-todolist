# grpc-todolist

## 运行
```bash
make env-up         # 启动环境
make user           # 启动用户摸块
make task           # 启动任务模块
make experimental   # 启动实验性模块
make api            # 启动网关
```

## 测试
### 接口
在`Postman`中导入`./postman`文件夹下的配置即可
需要编辑的变量有
- `host`: 网关地址
- `token`: 鉴权
### 负载均衡
这个项目支持轮询(Round-Robin)的负载均衡策略

只需要在`./config/config.yaml`中将`load-balance`设置为`true`便可启用该模块的负载均衡
```bash
make env-up api             # 开发环境与HTTP服务器
make experimental node=0    # rpc节点1
make experimental node=1    # rpc节点2
```
之后我们可以通过多次请求`experimental/ping`接口来测试负载均衡
## 实验性功能
在`./cmd`中，可以发现`experimental`的rpc服务端源代码，它包含以下方法
```proto
service ExperimentalService {
    rpc Ping                (Request)        returns (Response) {}          // Unary请求
    rpc ClientStream        (stream Request) returns (Response) {}          // 客户端发送流式请求
    rpc ServerStream        (Request)        returns (stream Response) {}   // 服务端回复流式请求
    rpc BidirectionalStream (stream Request) returns (stream Response) {}   // 双向流式请求
}
```
实验模块是一个简易版的rpc服务端，用来对grpc特性进行调试

我们可以在该实验性模块中探索grpc的stream特性，这部分代码逻辑在`./cmd/experimental/handler.go`中有较为完备的注释
## 项目结构

### 调用关系
```
                                      http
                           ┌────────────────────────┐
 ┌─────────────────────────┤                        ├───────────────────────────────┐
 │                         │          api           │                               │
 │      ┌─────────────────►|                        │◄──────────────────────┐       │
 │      │                  └───────────▲────────────┘                       │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                           resolve                                 │       │
 │      │                              │                                    │       │
req    resp                            │                                   resp    req
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                   ┌──────────▼─────────┐                          │       │
 │      │                   │                    │                          │       │
 │      │       ┌──────────►|        Etcd        |◄────────────────┐        │       │
 │      │       │           │                    │                 │        │       │
 │      │       │           └────────────────────┘                 │        │       │
 │      │       │                                                  │        │       │
 │      │     register                                           register   │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
┌▼──────┴───────┴───┐                                           ┌──┴────────┴───────▼─┐
│                   │───────────────── req ────────────────────►│                     │
│       user        │                                           │         task        │
│                   │◄──────────────── resp ────────────────────│                     │
└───────────────────┘                                           └─────────────────────┘
      protobuf                                                           protobuf

```
### 整体
```
.
├── Makefile
├── README.md
├── bin                 # 项目编译的二进制文件
├── cmd
│   ├── api             # 网关
│   ├── task            # 任务模块
│   ├── experimental    # 实验性功能
│   └── user            # 用户模块
├── config
│   ├── config.go
│   ├── config.yaml     # 配置项
│   └── sql             # sql建表语句
├── docker-compose.yml
├── go.mod
├── go.sum
├── postman             # 接口测试配置
├── idl
│   ├── pb
│   ├── task.proto
│   └── user.proto
└── pkg
    ├── errno           # 自定义错误
    └── utils           # 杂项
```
### 微服务
```
.
├── dal
│   ├── db              # 数据库操作
│   └── init.go
├── handler.go          # 请求解析/封装
├── main.go
├── model               # 数据库映射
│   └── model.go
├── pack                # 打包
│   └── pack.go
└── service             # 处理请求
    └── service.go
```
### 网关
```
.
├── handler             # 请求处理
│   ├── handler.go
│   ├── task.go
│   └── user.go
├── main.go
├── middleware          # gin中间件
│   └── jwt.go
├── routes              # 路由
│   └── routes.go
├── rpc                 # rpc调用
│   ├── init.go
│   ├── task.go
│   └── user.go
└── types               # 自定义类型
    ├── task.go
    ├── types.go
    └── user.go
```

## 指令列表
```bash
make env-up         # 启动环境
make env-down       # 结束环境
make proto          # 更新protoc
make user           # 启动用户摸块
make task           # 启动任务模块
make experimental   # 启动实验模块
make api            # 启动网关
```
### 多地址
在启动模块时，我们支持指定监听地址的序号，只需要在`config.yaml`中对模块设置多个监听地址，同时运行时添加`node`参数即可
```bash
make user node=0 # 运行user模块的第0个地址
```