package main

import (
	"context"
	"google.golang.org/grpc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
	timestamp "trip/proto/gen/timestamp"
)

// TimeStampUsage protobuf timestamp 类型示例
func TimeStampUsage(conn *grpc.ClientConn) {

	timestampClient := timestamp.NewTimestampEgClient(conn)
	timestampResp, err := timestampClient.AddTime(context.Background(), &timestamp.TimestampReq{
		Msg:     "timestamp 示例",
		AddTime: timestamppb.New(time.Now()),
	})
	if err != nil {
		log.Fatalf("cannot call timestamp.AddTime: %v\n", err)
	}
	log.Println("Resp.TraceId: ", timestampResp.TraceId)
}
