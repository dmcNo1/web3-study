package routers

import (
	"demo02/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitPermissionRouter(router *gin.Engine) {
	permissionRouter := router.Group("/permission")
	{
		db := models.DB
		permissionRouter.GET("/first", func(ctx *gin.Context) {
			permission := models.Permission{}
			db.Preload("Profile").First(&permission)
			ctx.JSON(http.StatusOK, permission)
		})

		permissionRouter.GET("/find", func(ctx *gin.Context) {
			permissions := []*models.Permission{}
			db.Preload("Roles").Find(&permissions)
			ctx.JSON(http.StatusOK, permissions)
		})
	}
}
