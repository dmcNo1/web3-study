package routers

import (
	"demo02/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitProfileRouter(router *gin.Engine) {
	profileRouter := router.Group("/profile")
	{
		db := models.DB
		profileRouter.GET("/first", func(ctx *gin.Context) {
			profile := models.Profile{}
			// Preload()可以关联查询指定的属性，不过一定要在model中配置好tag
			db.Preload("PermissionProfiles").Preload("Permission").First(&profile)
			ctx.JSON(http.StatusOK, profile)
		})
	}
}
