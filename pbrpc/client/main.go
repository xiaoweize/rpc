package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/xiaoweize/rpc/pbrpc/codec/client"
	"github.com/xiaoweize/rpc/pbrpc/service"
)

// 约束客户端
var _ service.HelloService = (*HelloServiceClient)(nil)

type HelloServiceClient struct {
	*rpc.Client
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	// 建立链接
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}
	// 采用Json编解码的客户端
	c := rpc.NewClientWithCodec(client.NewClientCodec(conn))
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request *service.Request, reply *service.Response) error {
	return p.Client.Call(service.HelloServiceName+".Hello", request, reply)
}

func main() {
	c, err := DialHelloService("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	var reply service.Response
	c.Hello(&service.Request{Value: "quyuan"}, &reply)
	fmt.Println(reply.Value)
}
