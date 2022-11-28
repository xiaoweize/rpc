package main

import (
	"context"
	"fmt"
	"net"

	"github.com/xiaoweize/rpc/grpc/server/pb"
	"google.golang.org/grpc"
)

type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

func (h *HelloService) Hello(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Value: fmt.Sprintf("hello,%s", req.Value)}, nil
}

func main() {
	//1.创建grpcserver 并注册服务端service
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &HelloService{})

	//2.创建socket连接
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
		return
	}
	//3.在tcp监听上启动grpc  内置了http2协议
	if err := grpcServer.Serve(listen); err != nil {
		fmt.Println(err)
		return
	}
}
