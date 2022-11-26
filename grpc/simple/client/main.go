package main

import (
	"context"
	"fmt"

	"github.com/xiaoweize/rpc/grpc/simple/server/pb"
	"google.golang.org/grpc"
)

func main() {

	//1.与grpc服务端建立网络连接，http2默认https 添加grpc.WithInsecure()参数
	conn, err := grpc.DialContext(context.Background(), "localhost:1234", grpc.WithInsecure())
	if err != nil {
		fmt.Println(conn)
		return
	}
	defer conn.Close()
	//2.在grpc连接上创建service client用于后面的本地方法调用
	//grpc protobuf已经生成了客户端调用服务端的SDK
	client := pb.NewHelloServiceClient(conn)

	//3.远程Hello方法调用
	resp, err := client.Hello(context.Background(), &pb.Request{Value: "quyuan"})
	if err != nil {
		fmt.Println(resp)
		return
	}
	fmt.Println(resp)
}
