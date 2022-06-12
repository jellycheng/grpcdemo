package clientdemo

import (
	"context"
	"fmt"
	"github.com/jellycheng/grpcdemo/pbservices/mygrpc/helloworld"
	"google.golang.org/grpc"
	"testing"
)

// go test -run=TestClient01
func TestClient01(t *testing.T)  {
	// 1. 创建连接对象
	conn, err := grpc.Dial(":9091", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(conn)
	// 2.实例化client对象
	client := helloworld.NewGoodHelloClient(conn)
	// 3. 处理入参数据
	req := &helloworld.HelloRequest{Ids: "1,4,7", Page: 1, PageSize: 15}
	// 4.发起请求，返回 HelloReply结构体指针对象
	response, err := client.SayHello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	// 响应结果: 接受到的入参数 ids=1,4,7:page=1:page_size=15
	fmt.Println(response.Message)
}
