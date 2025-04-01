package routers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitSessionRouter(router *gin.Engine) {
	sessionRouter := router.Group("/session")
	{
		sessionRouter.GET("/set", func(ctx *gin.Context) {
			// 获取session，如果session不存在，就会自动初始化一个session
			session := sessions.Default(ctx)
			// 设置过期时间
			session.Options(sessions.Options{
				MaxAge: 3600,
			})
			// 设置session
			session.Set("username", "jackpot")
			// 保存session
			session.Save()
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})

		sessionRouter.GET("/get", func(ctx *gin.Context) {
			sessions := sessions.Default(ctx)
			// 获取session中的数据
			username := sessions.Get("username")
			ctx.JSON(http.StatusOK, gin.H{"message": "success", "username": username})
		})
	}
}
