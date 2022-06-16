package clientdemo

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/jellycheng/grpcdemo/pbservices/mygrpc/helloworld"
	"github.com/jellycheng/grpcdemo/pbservices/mygrpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
	"time"
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

// go test -run=TestClient02
func TestClient02(t *testing.T)  {
	conn, err := grpc.Dial(":9091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(fmt.Sprintf("did not connect: %v", err))
	}
	defer conn.Close()
	c := helloworld.NewGoodHelloClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Ids: "9,8,7", Page: 3, PageSize: 30})
	if err != nil {
		fmt.Println(fmt.Sprintf("could not greet: %v", err))
	}
	fmt.Println(fmt.Sprintf("%s", r.GetMessage()))

}

// go test -run=TestClient03
func TestClient03(t *testing.T) {
	// 序列化
	userObj := user.UserInfoResp{Userid: 123, Username: "admin", Tags: []string{"高值用户", "电商用户"}}
	res, _ := proto.Marshal(&userObj)
	fmt.Println(string(res))
	// 反序列化
	userObjV2 := new(user.UserInfoResp)
	_ = proto.UnmarshalMerge(res, userObjV2)
	fmt.Println(userObjV2.Tags)
	fmt.Println(userObjV2.String())

}

