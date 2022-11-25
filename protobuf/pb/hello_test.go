package pb_test

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/xiaoweize/rpc/protobuf/pb"
)

func TestMarshal(t *testing.T) {
	should := assert.New(t)

	clientObj := &pb.String{Value: "hello"}
	//序列化
	pbBytes, err := proto.Marshal(clientObj)
	if should.NoError(err) {
		fmt.Println("marshal:", pbBytes)
	}

	ServerObj := &pb.String{}
	//反序列化
	err = proto.Unmarshal(pbBytes, ServerObj)
	if should.NoError(err) {
		fmt.Println("unmarshal:", ServerObj)
	}

}
