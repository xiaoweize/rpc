package main

import (
	"fmt"

	"github.com/xiaoweize/rpc/protobuf/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	clientObj := &pb.String{Value: "hello"}
	//序列化成byte
	out, err := proto.Marshal(clientObj)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("encode bytes:", out)
	//反序列化
	serverObj := &pb.String{}
	if err = proto.Unmarshal(out, serverObj); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("decode obj:", serverObj)
}
