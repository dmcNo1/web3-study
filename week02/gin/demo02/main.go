package main

import (
	"demo02/routers"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("main init")
}

func main() {
	router := gin.Default()

	routers.InitCookieRouter(router)

	// 配置session中间件
	// 创建基于Cookie的存储引擎，security888888是用于加密的秘钥
	store := cookie.NewStore([]byte("security888888"))
	// 配置session的中间件，store就是之前创建的存储引擎；当创建session的时候，会调用这个中间件
	router.Use(sessions.Sessions("mySession", store))
	routers.InitSessionRouter(router)

	routers.InitValidateRouter(router)

	router.Run(":8081")
}
