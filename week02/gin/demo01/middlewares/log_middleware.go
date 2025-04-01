package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 配置一个日志中间件
func LogMiddleware(ctx *gin.Context) {
	start := time.Now()
	// 给其他中间件传递参数
	ctx.Set("msgFromLog", "天音波")

	// 执行了Next之后，就会向后继续执行
	ctx.Next()
	// 知道所有的业务逻辑、后续中间件执行完，才会执行下面的逻辑
	costTime := time.Since(start)
	fmt.Printf("log --- 总耗时: %d\n", costTime)
}
