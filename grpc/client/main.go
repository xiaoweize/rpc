package main

import (
	"context"
	"fmt"

	"github.com/xiaoweize/rpc/grpc/server/pb"
	"google.golang.org/grpc"
)

func main() {
	//创建连接
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	//在连接上创建grpc客户端
	client := pb.NewHelloServiceClient(conn)
	//远程方法调用
	reply, err := client.Hello(context.Background(), &pb.Request{Value: "quyuan"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply)

}
