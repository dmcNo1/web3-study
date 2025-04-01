package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitCookieRouter(router *gin.Engine) {
	cookieRouter := router.Group("/cookie")
	{
		cookieRouter.GET("/", func(ctx *gin.Context) {
			// 获取cookie
			cookie, err := ctx.Cookie("key_cookie")
			if err != nil {
				cookie = "NotSet"
				// func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
				// name：key
				// value：value
				// maxAge：过期时间，单位为秒
				// path：有效路径，只有访问路径为/user才能访问这个Cookie
				// domain：域名，主要用于实现多个域名共享；比如两个网址，一个在192.168.100.1，一个在192.168.100.2，这两个都想使用就能这样
				// secure：当这个值为true时，只有在https中才能生效
				// httpOnly，微软对Cookie的扩展，防止XSS攻击
				ctx.SetCookie("key_cookie", cookie, 100, "/cookie", "localhost", false, true)
			}
			fmt.Printf("cookie = %s\n", cookie)
		})

		cookieRouter.GET("/home", AuthMiddleware, func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, map[string]string{"data": "success"})
		})
		cookieRouter.GET("/login", func(ctx *gin.Context) {
			auth, err := ctx.Cookie("auth_cookie")
			if err != nil || auth == "" {
				auth = "auth"
				ctx.SetCookie("auth_cookie", auth, 100, "/cookie", "localhost", false, true)
			}
			ctx.String(http.StatusOK, "success!")
		})
	}
}

// 模拟鉴权
func AuthMiddleware(ctx *gin.Context) {
	if auth, err := ctx.Cookie("auth_cookie"); err == nil {
		// 鉴权通过，放行，且无后续操作，所以需要手动return
		if auth == "auth" {
			ctx.Next()
			return
		}
	}

	// 鉴权失败，拦截，一旦执行了Abort，后续代码都会失效
	ctx.JSON(http.StatusUnauthorized, "please login!")
	ctx.Abort()
}
