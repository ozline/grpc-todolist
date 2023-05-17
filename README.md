# grpc-todolist

## 运行
对于任何模块，请在不同的Terminal中运行
```bash
make env-up     # 启动环境
make user       # 启动用户摸块
make task       # 启动任务模块
make api        # 启动网关模块
```

## 测试
在`Postman`中导入`./postman`文件夹下的配置即可
需要编辑的变量有
- `host`: 网关地址
- `token`: 鉴权

## 项目结构
### 整体
```
.
├── Makefile
├── README.md
├── bin                 # 项目编译的二进制文件
├── cmd
│   ├── api             # 网关
│   ├── task            # 任务模块
│   └── user            # 用户模块
├── config
│   ├── config.go
│   ├── config.yaml     # 配置项
│   └── sql             # sql建表语句
├── docker-compose.yml
├── go.mod
├── go.sum
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
我们可以在这个实验性模块中探索gRPC的Stream请求特性