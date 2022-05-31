package service

import (
	"fmt"
	"log"
	"sync"
	"time"
	trippb "trip/proto/gen/trip"
)

type StreamServer struct{}

// GetStream 服务器流模式
func (s *StreamServer) GetStream(req *trippb.StreamReqData, server trippb.StreamService_GetStreamServer) error {
	log.Println(req.In)
	i := 0
	for true {
		err := server.Send(&trippb.StreamRespData{
			Data: fmt.Sprintf("time: %v", time.Now().Unix()),
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
	return nil
}

// PutStream 客户端流模式
func (s *StreamServer) PutStream(server trippb.StreamService_PutStreamServer) error {
	for true {
		rev, err := server.Recv()
		if err != nil {
			log.Fatalf("Recieve client stream failed: %v", err)
		}
		log.Println(rev.In)
	}
	return nil
}

// AllStream 双向流模式
func (s *StreamServer) AllStream(server trippb.StreamService_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		i := 0
		for true {
			err := server.Send(&trippb.StreamRespData{
				Data: fmt.Sprintf("你好，我是服务器: %v", time.Now().Unix()),
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
			rev, err := server.Recv()
			if err != nil {
				log.Fatalf("Recieve client stream failed: %v", err)
			}
			log.Printf("收到客户端请求：%s", rev.In)
		}
	}()
	wg.Wait()
	return nil
}
