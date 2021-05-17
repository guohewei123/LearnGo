package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	rpcfunc "rpcdemo"
)


func main() {
	rpc.Register(rpcfunc.DemoService{})

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)

	}
}

// terminal 测试方法
// sh: telnet localhost 1234
// {"method": "DemoService.Div", "params": [{"A": 3, "B": 0}], "id": 123}
// 输出：{"id":123,"result":null,"error":"division by zero"}
// {"method": "DemoService.Div", "params": [{"A": 3, "B": 4}], "id": 123}
// 输出：{"id":123,"result":0.75,"error":null}
