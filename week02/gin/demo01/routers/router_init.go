package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter(router *gin.Engine) {
	initUserRouters(router)
	initAuthRouters(router)
	initMiddleWareRouter(router)
	initSyncRouter(router)

	// 配置404对应的返回结果
	router.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "404 ~~~~~")
	})
}
