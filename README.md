# grpcdemo
```
grpc研究、示例
```

## grpc环境
```
which protoc
which protoc-gen-go
查看版本： protoc --version

安装protoc编译器：
    Linux下安装示例： 
        sudo yum install -y protobuf protobuf-compiler protobuf-devel
    Mac下安装示例： brew install protobuf

安装go的proto、protoc-gen-go命令（用途：生成go代码插件）：
    go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
    ls $GOPATH/bin
    ls $(go env GOPATH)/bin
    设置环境变量： export PATH="$PATH:$(go env GOPATH)/bin"

安装grpc插件：
  git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
  或者 go get -u google.golang.org/grpc
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
  which protoc-gen-go-grpc
 
```

## grpc从0到1示例
```
cd grpcdemo
初始化mod:
    go mod init github.com/jellycheng/grpcdemo
创建./pbfiles/hello.proto文件
生成grpc代码：
    mkdir pbservices 
    protoc --go_out=plugins=grpc:./pbservices ./pbfiles/hello.proto
    或者 --proto_path 指定proto文件查找目录,多个目录用冒号分割，默认是protoc命令执行目录：
        protoc --proto_path=.:./pbfiles/ --go_out=plugins=grpc:./pbservices hello.proto
    文件名hello.proto 生成代码文件名为 hello.pb.go
    
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

