# grpcdemo
```
grpc研究、示例
```

## grpc环境
```
which protoc
which protoc-gen-go
查看版本： 
  protoc --version
  protoc-gen-go --version
  protoc-gen-go-grpc --version

安装protoc编译器：
    Linux下安装示例： 
        sudo yum install -y protobuf protobuf-compiler protobuf-devel
    Mac下安装示例： brew install protobuf

安装go的proto、protoc-gen-go命令（用途：生成go代码插件）：
  旧版本方式：
    go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
  新版本方式：  
    go install google.golang.org/protobuf/cmd/protoc-gen-go
    
    ls $GOPATH/bin
    ls $(go env GOPATH)/bin
    设置环境变量： export PATH="$PATH:$(go env GOPATH)/bin"
    which protoc-gen-go

安装grpc插件：
  git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
  或者 go get -u google.golang.org/grpc
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
  which protoc-gen-go-grpc
 

--go_out​​ 对应 ​​protoc-gen-go​​ 插件
​​--go-grpc_out​​ 对应 ​​protoc-gen-go-grpc​​ 插件
​​--*_out​​ 对应 ​​protoc-gen-*​​ 插件
​protoc-gen-go​​ 和 ​​protoc-gen-go-grpc​​ 这两个插件有什么不同：
  当使用参数 ​​--go_out=plugins=grpc:./xxx​​​ 生成时，生成的文件 ./xxx/​​*.pb.go​​​ 包含消息序列化代码和gRPC​​代码
  当使用参数 ​​--go_out=./xxx --go-grpc_out=./xxx​​​ 生成时，会生成两个文件 ​​./xxx/*.pb.go​​​ 和 ​​./xxx/*._grpc.pb.go​​​ ，分别是消息序列化代码和​​gRPC​​代码


```

## grpc从0到1示例
```
cd grpcdemo
初始化mod:
    go mod init github.com/jellycheng/grpcdemo
创建./pbfiles/hello.proto文件
生成grpc代码-旧版本命令：
    mkdir pbservices 
    protoc --go_out=plugins=grpc:./pbservices ./pbfiles/hello.proto
    文件名hello.proto 生成代码文件名为 hello.pb.go

生成grpc代码-新版本命令：
  此时hello.proto文件 option go_package = "mygrpc/helloworld"; 改成 option go_package = ".;mygrpc/helloworld";
  mkdir pbservices
  protoc --go_out=./pbservices --go_opt=paths=source_relative \
    --go-grpc_out=./pbservices --go-grpc_opt=paths=source_relative \
    ./pbfiles/hello.proto

实现暴露的服务方法： ./services/GoodHello.go
创建grpc服务： ./main.go
下载模块依赖： go mod tidy
启动grpc服务： go run main.go
grpc客户端调用服务示例： ./clientdemo/client_test.go

代码目录如下：
.
├── LICENSE
├── README.md
├── clientdemo
│   └── client_test.go
├── go.mod
├── go.sum
├── main.go
├── pbfiles
│   └── hello.proto
├── pbservices
│   └── mygrpc
│       └── helloworld
│           └── hello.pb.go
└── services
    └── GoodHello.go

```

## grpc资料
```
https://grpc.io/

```

