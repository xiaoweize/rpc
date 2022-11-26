package pb_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xiaoweize/rpc/protobuf/pb"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestXxx(t *testing.T) {
	should := assert.New(t)
	//any对象赋值
	es := &pb.ErrorStatus{Message: "test"}
	anyEs, err := anypb.New(es)
	if should.NoError(err) {
		fmt.Println(anyEs)
	}
	//any对象取值
	objReceive := pb.ErrorStatus{}
	err = anyEs.UnmarshalTo(&objReceive)
	if should.NoError(err) {
		fmt.Println(objReceive.Message)
	}

}
