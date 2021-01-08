package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/rpc"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	// 开启一个协程，监听 7000 端口，开启 rpc 服务
	go func() {
		// 1.注册服务
		rect := new(Rect)
		// 注册一个rect的服务
		rpc.Register(rect)
		// 2.服务处理绑定到http协议上
		rpc.HandleHTTP()
		// 3.监听服务
		err := http.ListenAndServe(":7000", nil)
		if err != nil {
			log.Panicln(err)
		}
	}()

	r.Run(":8080")
}

type Params struct {
	Width, Height int
}

type Rect struct{}

// RPC服务端方法，求矩形面积
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Height * p.Width
	return nil
}
