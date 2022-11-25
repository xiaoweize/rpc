package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/xiaoweize/rpc/json_tcp/service"
)

// 约束客户端
var _ service.HelloService = (*HelloServiceClient)(nil)

type HelloServiceClient struct {
	*rpc.Client
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	// c, err := rpc.Dial(network, address)
	// if err != nil {
	// 	return nil, err
	// }
	// 建立链接
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}
	// 采用Json编解码的客户端
	c := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(service.HelloServiceName+".Hello", request, reply)
}

func main() {
	c, err := DialHelloService("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	var reply string
	c.Hello("quyuan", &reply)
	fmt.Println(reply)
}
