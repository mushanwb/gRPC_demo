package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// 连接 rpc 服务
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	// 调用 rpc 注册的接口服务，并传入方法需要的参数
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
