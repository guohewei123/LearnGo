package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	trippb "trip/proto/gen/trip"
)

// GetTripClient 普通模式
func GetTripClient(conn *grpc.ClientConn) {
	tsClient := trippb.NewTripServiceClient(conn)

	r, err := tsClient.GetTrip(context.Background(), &trippb.GetTripRequest{
		Id: "trip123",
	})
	if err != nil {
		log.Fatalf("cannot call GetTrip: %v\n", err)
	}
	fmt.Println(r.Trip.Status.Number())
}

// GetTripClientWithMetadata metadata 示例
func GetTripClientWithMetadata(conn *grpc.ClientConn) {
	tsClient := trippb.NewTripServiceClient(conn)
	// 第一种方法
	//md := metadata.New(map[string]string{"key1": "value1", "key2": "value2"})

	// 第二种放过
	md := metadata.Pairs(
		"key1", "value1",
		"key1", "value",
		"key2", "value2")

	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := tsClient.GetTrip(ctx, &trippb.GetTripRequest{
		Id: "trip123[metadata 示例]",
	})
	if err != nil {
		log.Fatalf("cannot call GetTrip: %v\n", err)
	}
	fmt.Println("resp.Trip.Status=", r.Trip.Status.Number())
}


// AuthVerifyWithMetadata metadata 实现auth认证示例
func AuthVerifyWithMetadata(conn *grpc.ClientConn) {
	tsClient := trippb.NewTripServiceClient(conn)

	// 添加 auth Token
	md := metadata.Pairs("auth-token", "aaabbbccc")

	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := tsClient.GetTrip(ctx, &trippb.GetTripRequest{
		Id: "trip123[metadata 示例]",
	})
	if err != nil {
		log.Fatalf("cannot call GetTrip: %v\n", err)
	}
	fmt.Println("resp.Trip.Status=", r.Trip.Status.Number())
}


// AuthVerifyWithMetadata1 metadata 实现auth认证示例
func AuthVerifyWithMetadata1(conn *grpc.ClientConn) {
	tsClient := trippb.NewTripServiceClient(conn)

	// 添加 auth Token
	md := metadata.Pairs("auth-token", "aaabbbccc")

	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := tsClient.GetTrip(ctx, &trippb.GetTripRequest{
		Id: "trip123[metadata 示例]",
	})
	if err != nil {
		log.Fatalf("cannot call GetTrip: %v\n", err)
	}
	fmt.Println("resp.Trip.Status=", r.Trip.Status.Number())
}