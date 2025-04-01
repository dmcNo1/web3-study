package routers

import (
	"demo01/controller/user"

	"github.com/gin-gonic/gin"
)

func initUserRouters(router *gin.Engine) {
	// 创建一个路由分组，所有/user的路由都会匹配到这里头
	userRouters := router.Group("/user")
	// {} 是书写规范
	{
		// 绑定路由到对应的controller，不同的项目不一样，很多项目都不是mvc架构的
		userController := user.UserController{}
		userRouters.GET("/user/:name/*action", userController.Action)
		userRouters.POST("/form", userController.Form)
		userRouters.POST("/upload", userController.Upload)
		userRouters.POST("/uploadFiles", userController.UploadFiles)
	}
}
