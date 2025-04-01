package main

import (
	"demo01/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 配置一个默认的路由
	router := gin.Default()

	// 设置最大上传文件大小
	router.MaxMultipartMemory = 8 << 20

	// 绑定路由
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World!")
	})

	// 初始化路由
	routers.InitRouter(router)

	// 启动一个Web服务，可以通过router.Run([address]:port)指定端口，router.Run("127.0.0.1:8081")
	router.Run(":8081")
}
