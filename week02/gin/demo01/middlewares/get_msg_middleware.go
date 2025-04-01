package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetMsgMiddleware(ctx *gin.Context) {
	// 获取其他中间件存入的数据
	if msg, existsFlag := ctx.Get("msgFromLog"); existsFlag {
		fmt.Printf("msgFromLog = %v\n", msg)
	}
	ctx.Next()
}
