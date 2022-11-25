package main

import (
	"net"
	"net/rpc"

	"github.com/xiaoweize/rpc/pbrpc/codec/server"
	"github.com/xiaoweize/rpc/pbrpc/service"
)

// 通过接口约束HelloService服务
var _ service.HelloService = (*HelloService)(nil)

//hello handler
type HelloService struct{}

//固定格式函数
func (h *HelloService) Hello(req *service.Request, resp *service.Response) error {
	resp.Value = "hello:" + req.Value
	return nil
}

func main() {
	// 把我们的对象注册成一个rpc的 receiver
	// 其中rpc.Register函数调用会将对象类型中所有满足RPC规则的对象方法注册为RPC函数，
	// 所有注册的方法会放在“HelloService”服务空间之下
	rpc.RegisterName("HelloService", new(HelloService))
	// 然后我们建立一个唯一的TCP链接，
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	// 通过rpc.ServeConn函数在该TCP链接上为对方提供RPC服务。
	for {
		//Accept接受listener上的连接。Accept会阻塞,直到接收到连接，并返回这个conn。
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		// go rpc.ServeConn(conn)
		// 服务端代码中最大的变化是用rpc.ServeCodec函数替代了rpc.ServeConn函数，
		// 传入的参数是针对服务端的json编解码器
		go rpc.ServeCodec(server.NewServerCodec(conn))
	}

}
