package routers

import (
	"demo01/controller/auth"

	"github.com/gin-gonic/gin"
)

// 这里懒得写前端了，自己用postman玩
func initAuthRouters(router *gin.Engine) {
	authRouter := router.Group("/auth")
	{
		authController := auth.AuthController{}
		authRouter.POST("/loginJson", authController.LoginJson)
		authRouter.POST("/loginForm", authController.LoginForm)
		authRouter.GET("/loginUri/:username/:password", authController.LoginUri)
	}
}
