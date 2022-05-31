package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	mapbp "trip/proto/gen/map"
)

// MapUsage protobuf map 类型示例
func MapUsage(conn *grpc.ClientConn) {

	mapClient := mapbp.NewMapClient(conn)
	mapResp, err := mapClient.MapUsage(context.Background(), &mapbp.MapReq{
		Msg:     "protobuf map 类型示例",
		MapData: map[string]string{"name": "张三", "url": "https://192.168.10.33/map"},
	})
	if err != nil {
		log.Fatalf("cannot call saymap: %v\n", err)
	}
	log.Println("Resp.TraceId: ", mapResp.TraceId)
}
