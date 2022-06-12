package services

import (
	"context"
	"fmt"
	"github.com/jellycheng/grpcdemo/pbservices/mygrpc/helloworld"
)

// GoodHelloImpl 实现服务
type GoodHelloImpl struct {

}

// SayHello 实现服务的方法
func (m *GoodHelloImpl)SayHello(ctx context.Context, reqDto *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	// 业务逻辑封装todo
	ret := &helloworld.HelloReply{
		Message: fmt.Sprintf("响应结果: 接受到的入参数 ids=%s:page=%d:page_size=%d",
										reqDto.Ids,
										reqDto.Page,
										reqDto.PageSize),
	}
	return ret, nil
}
