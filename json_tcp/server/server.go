package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/xiaoweize/rpc/json_tcp/service"
)

// 通过接口约束HelloService服务
var _ service.HelloService = (*HelloService)(nil)

//hello handler
type HelloService struct{}

//固定格式函数
func (h *HelloService) Hello(req string, resp *string) error {
	*resp = "hello:" + req
	return nil
}

func NewRPCReadWriteCloserFromHTTP(w http.ResponseWriter, r *http.Request) *RPCReadWriteCloser {
	return &RPCReadWriteCloser{w, r.Body}
}

type RPCReadWriteCloser struct {
	io.Writer
	io.ReadCloser
}

func main() {
	// 把我们的对象注册成一个rpc的 receiver
	// 其中rpc.Register函数调用会将对象类型中所有满足RPC规则的对象方法注册为RPC函数，
	// 所有注册的方法会放在“HelloService”服务空间之下
	rpc.RegisterName("HelloService", new(HelloService))

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		rpc.ServeCodec(jsonrpc.NewServerCodec(NewRPCReadWriteCloserFromHTTP(w, r)))
	})
	http.ListenAndServe(":1234", nil)

}
