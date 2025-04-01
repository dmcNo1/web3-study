package routers

import (
	"demo01/controller/middleware"
	"demo01/middlewares"

	"github.com/gin-gonic/gin"
)

func initMiddleWareRouter(router *gin.Engine) {
	middleWareRouter := router.Group("/middleware")
	{
		// 使用一个全局中间件，在这个路由下，所有的请求都会触发这个中间件；中间件是按照配置顺序执行的，全局中间件配置的最早，所以会先执行LogMiddleware
		// Use方法定义：
		// 	func (group *RouterGroup) Use(middleware ...HandlerFunc) IRoutes
		// 	type HandlerFunc func(*Context)
		// 所以只需要使用一个函数传参为*Context的即可作为中间件
		middleWareRouter.Use(middlewares.LogMiddleware)
		// 局部中间件，GetMsgMiddleware在LogMiddleware之后配置，所以会在LogMiddleware之后执行
		middleWareRouter.GET("/get_msg", middlewares.GetMsgMiddleware, middleware.GetMsg)
	}
}
