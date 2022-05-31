package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
	rpcfunc "rpcdemo"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)
	var result float64
	err = client.Call("DemoService.Div", rpcfunc.Args{10, 3}, &result)
	if err != nil {
		log.Printf("err: %s\n", err)
	} else {
		fmt.Println(result)
	}

	err = client.Call("DemoService.Div", rpcfunc.Args{10, 0}, &result)
	if err != nil {
		log.Printf("err: %s\n", err)
	} else {
		fmt.Println(result)
	}
}
