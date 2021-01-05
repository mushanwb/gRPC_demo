package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {
	// 注册 rpc 接口服务
	rpc.RegisterName("HelloService", new(HelloService))

	// 服务监听的端口 1234， tcp 连接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	// 时时监听端口事件
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}

// 数据类型
type HelloService struct{}

// 对外提供的 rpc 服务
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}
