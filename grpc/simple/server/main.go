package main

import (
	"context"
	"fmt"
	"net"

	"github.com/xiaoweize/rpc/grpc/simple/server/pb"
	"google.golang.org/grpc"
)

var _ pb.HelloServiceServer = (*HelloService)(nil)

// UnimplementedHelloServiceServer must be embedded（嵌入） to have forward compatible implementations.
type HelloService struct {
	//用于实现HelloServiceServer接口，注意要重新定义下hello方法，直接使用会抛出"method Hello not implemented"错误
	pb.UnimplementedHelloServiceServer
}

func (h *HelloService) Hello(c context.Context, req *pb.Request) (*pb.Response, error) {
	resp := &pb.Response{}
	resp.Value = "hello:" + req.Value

	return resp, nil
}

func main() {
	//1.注册本地service
	// 首先通过grpc.NewServer()构造一个gRPC服务对象
	grpcServer := grpc.NewServer()
	// 然后通过gRPC插件生成的RegisterHelloServiceServer函数注册我们实现的HelloServiceServer服务
	pb.RegisterHelloServiceServer(grpcServer, &HelloService{})

	//2.打开本地监听
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	//3.在监听的端口上提供grpc服务 内置http2
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println(err)
		return
	}

}
