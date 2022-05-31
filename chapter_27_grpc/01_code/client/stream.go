package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
	trippb "trip/proto/gen/trip"
)

// GetStreamClient 服务器 发送流 到客户端
func GetStreamClient(conn *grpc.ClientConn) {
	streamClient := trippb.NewStreamServiceClient(conn)

	stream, err := streamClient.GetStream(context.Background(), &trippb.StreamReqData{
		In: "给我来个数据流",
	})
	if err != nil {
		log.Fatalf("cannot call ServerStream: %v", err)
	}

	for true {
		rev, err := stream.Recv()
		if err != nil {
			log.Fatalf("receive stream err: %v", err)
		}
		fmt.Println(rev.Data)
	}
}

// PutStreamClient 客户端 发送流 到服务端
func PutStreamClient(conn *grpc.ClientConn) {
	streamClient := trippb.NewStreamServiceClient(conn)
	stream, err := streamClient.PutStream(context.Background())
	if err != nil {
		log.Fatalf("cannot call PutStream: %v", err)
	}
	i := 0
	for true {
		i++
		err := stream.Send(&trippb.StreamReqData{
			In: fmt.Sprintf("我给你推个流啊：%d", i),
		})
		if err != nil {
			log.Fatalf("Send stream to Server Failed: %v", err)
		}
		time.Sleep(time.Second)
		if i > 10 {
			fmt.Println("不推啦，拜拜")
			err := stream.CloseSend()
			if err != nil {
				log.Fatalf("关闭发送通道失败啦: %v", err)
			}
			break
		}
	}
}

// AllStreamClient 客户端 和 服务端 同时发送流
func AllStreamClient(conn *grpc.ClientConn) {
	streamClient := trippb.NewStreamServiceClient(conn)
	stream, err := streamClient.AllStream(context.Background())
	if err != nil {
		log.Fatalf("cannot call PutStream: %v", err)
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		i := 0
		for true {
			err := stream.Send(&trippb.StreamReqData{
				In: fmt.Sprintf("你好，我是客户端: %v", time.Now().Unix()),
			})
			if err != nil {
				log.Fatalf("ServerStream Send Message Failed: %v", err)
			}
			time.Sleep(time.Second)
			i++
			if i > 10 {
				log.Println("发送结束啦!!")
				break
			}
		}
	}()

	go func() {
		defer wg.Done()
		for true {
			rev, err := stream.Recv()
			if err != nil {
				log.Fatalf("Recieve client stream failed: %v", err)
			}
			log.Printf("收到服务器数据: %s", rev.Data)
		}
	}()
	wg.Wait()
}
