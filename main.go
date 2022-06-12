package main

import (
	"fmt"
	"github.com/jellycheng/grpcdemo/pbservices/mygrpc/helloworld"
	"github.com/jellycheng/grpcdemo/services"
	"google.golang.org/grpc"
	"net"
)

func main()  {
	// 1.获取grpc对象,返回grpc.Server结构体指针
	rpcserver := grpc.NewServer()
	// 2.注册服务
	helloworld.RegisterGoodHelloServer(rpcserver, new(services.GoodHelloImpl))

	// 3.启动服务
	lis, err := net.Listen("tcp", ":9091") // 设置监听端口
	if err != nil {
		fmt.Println(err)
		return
	}
	err = rpcserver.Serve(lis)
	if err != nil {
		fmt.Println(err)
	}
}
