package sync

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// 同步
func Async(ctx *gin.Context) {
	time.Sleep(time.Second * 3)
	log.Printf("同步执行：%s\n", ctx.Request.URL.Path)
}

// 异步
func Sync(ctx *gin.Context) {
	// 创建一个副本，不然原上下文可能会被回收
	copyCtx := ctx.Copy()
	go func() {
		time.Sleep(time.Second * 3)
		log.Printf("异步执行：%s\n", copyCtx.Request.URL.Path)
	}()
}
