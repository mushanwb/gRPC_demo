package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/rpc"
)

// 传的参数
type Params struct {
	Width, Height int
}

// rpc 连接实例
var rpcConn *rpc.Client

// 连接 rpc 服务的 7000 端口
func ConnRpc() {
	var err error
	rpcConn, err = rpc.DialHTTP("tcp", ":7000")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := gin.Default()

	ConnRpc()

	r.GET("/pong", func(c *gin.Context) {
		ret := 0
		// 开始调用 rpc 注册的服务，以及调用服务方法
		err2 := rpcConn.Call("Rect.Area", Params{50, 100}, &ret)
		if err2 != nil {
			log.Fatal(err2)
		}
		c.JSON(200, gin.H{
			"message": ret,
		})
	})

	r.Run(":8081")
}
