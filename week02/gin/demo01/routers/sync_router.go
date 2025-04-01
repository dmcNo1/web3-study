package routers

import (
	"demo01/controller/sync"

	"github.com/gin-gonic/gin"
)

func initSyncRouter(router *gin.Engine) {
	syncRouter := router.Group("/sync")
	{
		syncRouter.GET("/async", sync.Async)
		syncRouter.GET("/sync", sync.Sync)
	}
}
