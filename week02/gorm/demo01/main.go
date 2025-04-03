package main

import (
	"demo01/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.InitUserRouters(router)
	routers.InitCreditCardRouters(router)
	router.Run(":8081")
}
