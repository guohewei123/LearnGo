package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, calculateReqTimeConsumingOption())
	opts = append(opts, addRPCCredentialsOption())

	conn, err := grpc.Dial("localhost:8081", opts...)
	if err != nil {
		log.Fatalf("cannot connect server: %v", err)
	}
	defer conn.Close()

	//TripServiceClient
	//GetTripClient(conn)
	//GetStreamClient(conn)
	//PutStreamClient(conn)
	//AllStreamClient(conn)
	//NestedMessageUsage(conn)
	//MapUsage(conn)
	//TimeStampUsage(conn)
	//GetTripClientWithMetadata(conn)
	//AuthVerifyWithMetadata(conn)
	//ErrorUsage(conn)
	ErrorUsageTimeOut(conn)
}


// 1. grpc 拦截器 实现 计算请求耗时
func calculateReqTimeConsumingOption() grpc.DialOption {

	clientInterceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		log.Println("开始发起请求..............")
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		log.Printf("请求结束!!!--------耗时：%s\n", time.Since(start))
		return err
	}
	return grpc.WithUnaryInterceptor(clientInterceptor)
}

// 2. 通过 客户端拦截器 grpc.WithPerRPCCredentials 和服务端 metadata + 拦截器 实现 auth 认证
type grpcCredential struct{}

func (g grpcCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"auth-token": "aaabbbccc"}, nil
}

func (g grpcCredential) RequireTransportSecurity() bool {
	return false
}

func addRPCCredentialsOption() grpc.DialOption {
	return grpc.WithPerRPCCredentials(grpcCredential{})
}
