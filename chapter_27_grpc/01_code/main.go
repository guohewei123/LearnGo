package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	"time"
	errortestpb "trip/proto/gen/errortest"
	mapbp "trip/proto/gen/map"
	nestedMsg "trip/proto/gen/nested_message"
	timestampbp "trip/proto/gen/timestamp"
	trippb "trip/proto/gen/trip"
	"trip/service"
)

func main() {
	log.SetFlags(log.Lshortfile)
	go startGRPCGateWay()
	RunGRPCServer()
}

// RunGRPCServer 启动GRPC server
func RunGRPCServer() {
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

	}
	var opts []grpc.ServerOption
	opts = append(opts, UnaryServerOption())
	s := grpc.NewServer(opts...)
	registerServer(s)
	log.Fatal(s.Serve(listen))
}

// 组成 GPRC 服务
func registerServer(s *grpc.Server) {
	trippb.RegisterTripServiceServer(s, &service.TripService{})
	trippb.RegisterStreamServiceServer(s, &service.StreamServer{})
	nestedMsg.RegisterHelloServer(s, &service.NestedService{})
	mapbp.RegisterMapServer(s, &service.MapService{})
	timestampbp.RegisterTimestampEgServer(s, &service.TimestampService{})
	errortestpb.RegisterErrorTestServer(s, &service.ErrorService{})
}

// UnaryServerOption 拦截器实现 1. 拦截器验证 auth-token 2.请求处理耗时
// 注意 一元服务器拦截器(grpc.UnaryInterceptor) 只能设置一次
func UnaryServerOption() grpc.ServerOption {

	verifyAuth := func(ctx context.Context) error {
		fmt.Println("开始验证 auth token")
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if token, ok := md["auth-token"]; ok {
				if token[0] == "aaabbbccc" {
					return nil
				}
				return status.Error(codes.Unauthenticated, "token错误")
			}
			return status.Error(codes.Unauthenticated, "token认证信息")
		} else {
			// 使用GRPC Error
			return status.Error(codes.Unauthenticated, "无token认证信息")
		}
	}

	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("接收到了一个新的请求...........")

		// 1. 拦截器验证 auth token
		err = verifyAuth(ctx)
		if err != nil {
			log.Printf("Auth Failed: %v\n", err)
			//return resp, err
		}

		// 2. 拦截器实现计算请求处理耗时
		start := time.Now()
		res, err := handler(ctx, req)
		fmt.Printf("请求处理完成!!!--------耗时: %s\n", time.Since(start))
		if err != nil {
			log.Printf("Exec hander error: %v\n", err)
		}
		return res, err
	}
	return grpc.UnaryInterceptor(interceptor)
}

// startGRPCGateWay 启动 GRPC GateWay
func startGRPCGateWay() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{
			EnumsAsInts: true, // 枚举类型输出整形 "不然会输出枚举key字符串"
			OrigName:    true, // 使用原始名字与 protobuf 文件定义字段一致
		}))
	err := trippb.RegisterTripServiceHandlerFromEndpoint(
		c,
		mux,
		"localhost:8081",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatalf("cannot start grpc gateway: %v", err)
	}
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("cannot listen and server: %v", err)
	}
}
