package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"
	errortestpb "trip/proto/gen/errortest"
)

// ErrorUsage grpc 返回error示例
func ErrorUsage(conn *grpc.ClientConn) {

	errorClient := errortestpb.NewErrorTestClient(conn)
	errorResp, err := errorClient.ErrorUsage(context.Background(), &emptypb.Empty{})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			log.Printf("Error code: %s, Error message: %s\n", st.Code(), st.Message())
			return
		}
		log.Printf("Call ErrorUsage return unkown error: %v\n", err)
	}
	log.Println("Resp.TraceId: ", errorResp.TraceId)
}

// ErrorUsageTimeOut grpc 设置超时返回error示例
func ErrorUsageTimeOut(conn *grpc.ClientConn) {
	errorClient := errortestpb.NewErrorTestClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	errorResp, err := errorClient.ErrorUsage(ctx, &emptypb.Empty{})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			log.Printf("Error code: %s, Error message: %s\n", st.Code(), st.Message())
			return
		}
		log.Printf("Call ErrorUsage return unkown error: %v\n", err)
	}
	log.Println("Resp.TraceId: ", errorResp.TraceId)
}
