package routers

import (
	"demo01/models"

	"github.com/gin-gonic/gin"
)

func InitCreditCardRouters(router *gin.Engine) {
	db := models.DB
	creditCardRouter := router.Group("/credit_card")
	{
		creditCardRouter.GET("/create", func(ctx *gin.Context) {
			db.AutoMigrate(&models.CreditCard{})
		})
	}
}
