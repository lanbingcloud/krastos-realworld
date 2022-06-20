本次演示的内容主要包含以下几点
- 安装kratos 
- 项目创建和介绍
- 依赖注入
- API接口
- 中间件


## 安装kratos (推荐在linux/mac环境下进行开发)
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## 项目创建和介绍

### 创建

```
# Create a template project
kratos new kratos-realworld
```
### 介绍

```
  .
├── Dockerfile  
├── LICENSE
├── Makefile  
├── README.md
├── api // 下面维护了微服务使用的proto文件以及根据它们所生成的go文件
│   └── helloworld
│       └── v1
│           ├── error_reason.pb.go
│           ├── error_reason.proto
│           ├── error_reason.swagger.json
│           ├── greeter.pb.go
│           ├── greeter.proto
│           ├── greeter.swagger.json
│           ├── greeter_grpc.pb.go
│           └── greeter_http.pb.go
├── cmd  // 整个项目启动的入口文件
│   └── server
│       ├── main.go
│       ├── wire.go  // 我们使用wire来维护依赖注入
│       └── wire_gen.go
├── configs  // 这里通常维护一些本地调试用的样例配置文件
│   └── config.yaml
├── generate.go
├── go.mod
├── go.sum
├── internal  // 该服务所有不对外暴露的代码，通常的业务逻辑都在这下面，使用internal避免错误引用
│   ├── biz   // 业务逻辑的组装层，类似 DDD 的 domain 层，data 类似 DDD 的 repo，而 repo 接口在这里定义，使用依赖倒置的原则。
│   │   ├── README.md
│   │   ├── biz.go
│   │   └── greeter.go
│   ├── conf  // 内部使用的config的结构定义，使用proto格式生成
│   │   ├── conf.pb.go
│   │   └── conf.proto
│   ├── data  // 业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口。我们可能会把 data 与 dao 混淆在一起，data 偏重业务的含义，它所要做的是将领域对象重新拿出来，我们去掉了 DDD 的 infra层。
│   │   ├── README.md
│   │   ├── data.go
│   │   └── greeter.go
│   ├── server  // http和grpc实例的创建和配置
│   │   ├── grpc.go
│   │   ├── http.go
│   │   └── server.go
│   └── service  // 实现了 api 定义的服务层，类似 DDD 的 application 层，处理 DTO 到 biz 领域实体的转换(DTO -> DO)，同时协同各类 biz 交互，但是不应处理复杂逻辑
│       ├── README.md
│       ├── greeter.go
│       └── service.go
└── third_party  // api 依赖的第三方proto
    ├── README.md
    ├── google
    │   └── api
    │       ├── annotations.proto
    │       ├── http.proto
    │       └── httpbody.proto
    └── validate
        ├── README.md
        └── validate.proto
```

## 依赖注入

[wire](https://github.com/google/wire)是由 google 开源的一个供 Go 语言使用的依赖注入代码生成工具。它能够根据你的代码，生成相应的依赖注入 go 代码。

**Wire** 是一个灵活的依赖注入工具，通过自动生成代码的方式在编译期完成依赖注入。

在各个组件之间的依赖关系中，通常鼓励显式初始化，而不是全局变量传递。

所以通过 *Wire* 进行初始化代码，可以很好地解决组件之间的耦合，以及提高代码维护性。

### 安装

```
# 导入到项目中
go get -u github.com/google/wire

# 安装命令
go install github.com/google/wire/cmd/wire
```

### 原理

Wire 具有两个基本概念：Provider 和 Injector
Provider 是一个普通的 Go Func ，这个方法也可以接收其它 Provider 的返回值，从而形成了依赖注入；

```go
// 提供一个配置文件（也可能是配置文件）
func NewConfig() *conf.Data {...}

// 提供数据组件，依赖了数据配置（初始化 Database、Cache 等）
func NewData(c *conf.Data) (*Data, error) {...}

// 提供持久化组件，依赖数据组件（实现 CURD 持久化层）
func NewUserRepo(d *data.Data) (*UserRepo, error) {...}
```

### 使用

![](https://go-kratos.dev/images/wire.png)在每个模块中，只需要一个 *ProviderSet* 提供者集合，就可以在 wire 中进行依赖注入；

并且我们在每个组件提供入口即可，不需要其它依赖，例如：

```
-data
--data.go    // var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)
--greeter.go // func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {...}
```

### 初始化组件

```go
// 应用程序入口
cmd
-main.go
-wire.go
-wire_gen.go

// main.go 创建 kratos 应用生命周期管理
func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, greeter *service.GreeterService) *kratos.App {
    pb.RegisterGreeterServer(gs, greeter)
    pb.RegisterGreeterHTTPServer(hs, greeter)
    return kratos.New(
        kratos.Name(Name),
        kratos.Version(Version),
        kratos.Logger(logger),
        kratos.Server(
            hs,
            gs,
        ),
    )
}

// wire.go 初始化模块
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, error) {
    // 构建所有模块中的 ProviderSet，用于生成 wire_gen.go 自动依赖注入文件
    panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
```

生成代码，往`Makefile`添加以下代码

```makefile
# wire
.PHONY: wire
wire: 
	cd cmd/realworld/ && wire
```

## API接口

### 接口定义

通过 Protobuf IDL 定义对应的 REST API 和 gRPC API：

api/helloworld/v1/greeter.proto

```protobuf
syntax = "proto3";

package helloworld.v1;

import "google/api/annotations.proto";

option go_package = "github.com/go-kratos/service-layout/api/helloworld/v1;v1";
option java_multiple_files = true;
option java_package = "dev.kratos.api.helloworld.v1";
option java_outer_classname = "HelloWorldProtoV1";

// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply)  {
    option (google.api.http) = {
        // 定义一个 GET 接口，并且把 name 映射到 HelloRequest
        get: "/helloworld/{name}",
        // 可以添加附加接口
        additional_bindings {
            // 定义一个 POST 接口，并且把 body 映射到 HelloRequest
            post: "/v1/greeter/say_hello",
            body: "*",
        }
    };
  }
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```

### 生成接口

```sh
# 生成 proto 模板
kratos proto add api/helloworld/v1/greeter.proto
# 生成 client 源码
kratos proto client api/helloworld/v1/greeter.proto
# 生成 server 源码
kratos proto server api/helloworld/v1/greeter.proto -t internal/service
```



```
client:
|____api
| |____helloworld
| | |____v1
| | | |____greeter.pb.go
| | | |____greeter.proto
| | | |____greeter_http.pb.go
| | | |____greeter_grpc.pb.go

server:
| |____service
| | |____greeter.go
```

### 注册接口

**HTTP API** 是通过 protoc-gen-go-http 插件进行生成 http.Handler，然后可以注册到 HTTP Server 中：

```go
import "github.com/go-kratos/kratos/v2/transport/http"

greeter := &GreeterService{}
srv := http.NewServer(http.Address(":8000"))
srv.HandlePrefix("/", v1.NewGreeterHandler(greeter))
```

**gRPC API** 是通过 protoc-gen-go-grpc 插件进行生成 gRPC Register，然后可以注册到 GRPC Server 中:

```go
import "github.com/go-kratos/kratos/v2/transport/grpc"

greeter := &GreeterService{}
srv := grpc.NewServer(grpc.Address(":9000"))
v1.RegisterGreeterServer(srv, greeter)
```

### API文档

OpenAPI Swagger 使用

安装

```go
go get -u github.com/go-kratos/swagger-api
```

然后在`internal/server/http.go`的NewHTTPServer中进行初始化和注册，请尽量将这个路由注册放在最前面，以免匹配不到

```go
import "github.com/go-kratos/swagger-api/openapiv2"

openAPIhandler := openapiv2.NewHandler()
srv.HandlePrefix("/q/", openAPIhandler)
```

## 中间件接入

概念

中间件用于处理通用场景，支持自定义中间件

内置的中间件

- logging: 用于请求日志的记录。
- metrics: 用于启用metric。
- recovery: 用于recovery panic。
- tracing: 用于启用trace。
- validate: 用于处理参数校验。
- metadata: 用于启用元信息传递
- auth: 用于提供基于JWT的认证请求
- ratelimit: 用于服务端流量限制
- circuitbreaker: 用于客户端熔断控制

### 生效顺序

一个请求进入时的处理顺序为Middleware注册的顺序，而响应返回的处理顺序为注册顺序的倒序。

```
         ┌───────────────────┐
         │MIDDLEWARE 1       │
         │ ┌────────────────┐│
         │ │MIDDLEWARE 2    ││
         │ │ ┌─────────────┐││
         │ │ │MIDDLEWARE 3 │││
         │ │ │ ┌─────────┐ │││
REQUEST  │ │ │ │  YOUR   │ │││  RESPONSE
   ──────┼─┼─┼─▷ HANDLER ○─┼┼┼───▷
         │ │ │ └─────────┘ │││
         │ │ └─────────────┘││
         │ └────────────────┘│
         └───────────────────┘
```

### 使用中间件

```go
// http
// 定义opts
var opts = []http.ServerOption{
    http.Middleware(
        recovery.Recovery(), // 把middleware按照需要的顺序加入
        tracing.Server(),
        logging.Server(),
    ),
}
// 创建server
http.NewServer(opts...)



//grpc
var opts = []grpc.ServerOption{
    grpc.Middleware(
        recovery.Recovery(),  // 把middleware按照需要的顺序加入
        tracing.Server(),
        logging.Server(),
    ),
}
// 创建server
grpc.NewServer(opts...)
```

### 自定义中间件

需要实现`Middleware`接口。 中间件中您可以使用`tr, ok := transport.FromServerContext(ctx)`获得Transporter实例以便访问接口相关的元信息

基本的代码模板

```go
import (
    "context"

    "github.com/go-kratos/kratos/v2/middleware"
    "github.com/go-kratos/kratos/v2/transport"
)

func Middleware1() middleware.Middleware {
    return func(handler middleware.Handler) middleware.Handler {
        return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
            if tr, ok := transport.FromServerContext(ctx); ok {
                // Do something on entering 
                defer func() { 
                // Do something on exiting
                 }()
            }
            return handler(ctx, req)
        }
    }
}
```

