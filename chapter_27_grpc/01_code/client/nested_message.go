package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	nestedMsg "trip/proto/gen/nested_message"
)

// NestedMessageUsage protobuf 嵌套 message 示例
func NestedMessageUsage(conn *grpc.ClientConn) {

	// 如果需要单独使用 嵌套message
	data := nestedMsg.HelloReq_Result{
		Name: "大漂亮",
		Url:  "https://dapiaoliang.jj.com",
	}
	log.Printf("单独使用 嵌套message: %v\n", &data)

	helloClient := nestedMsg.NewHelloClient(conn)
	helloResp, err := helloClient.SayHello(context.Background(), &nestedMsg.HelloReq{
		Msg: "你好，祖国！",
		Data: []*nestedMsg.HelloReq_Result{
			&data,
		},
	})
	if err != nil {
		log.Fatalf("cannot call sayHello: %v\n", err)
	}
	log.Println("Resp.TraceId: ", helloResp.TraceId)
}


