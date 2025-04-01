package routers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Age      int       `form:"age" binding:"required,gt=10"` // age必须不为空，且大于10
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	Name     string
}

func InitValidateRouter(ctx *gin.Engine) {
	validateRouter := ctx.Group("/validate")
	{
		// 结构体验证，http://localhost:8081/validate/struct?age=100&birthday=2006-01-02
		validateRouter.GET("/struct", func(ctx *gin.Context) {
			var user User
			if err := ctx.ShouldBind(&user); err != nil {
				ctx.String(http.StatusInternalServerError, "error data")
				return
			}
			ctx.String(http.StatusOK, fmt.Sprintf("%#v", user))
		})
	}
}
