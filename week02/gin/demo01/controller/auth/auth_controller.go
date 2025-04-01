package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 传参对应的结构体，对应Param
// 可以在"“"之间，定义不同格式下该字段对应的参数，也可以配置是否必填
type Login struct {
	Username string `form:"un" json:"username" uri:"username" xml:"username" binding:"required"`
	Pwd      string `form:"pwd" json:"password" uri:"password" xml:"password" binding:"required"`
}

type AuthController struct{}

func (a *AuthController) LoginJson(ctx *gin.Context) {
	var login Login
	// 按照JSON格式绑定到对象
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, login)
}

func (a *AuthController) LoginForm(ctx *gin.Context) {
	var login Login
	// 按照JSON格式绑定到对象
	if err := ctx.ShouldBind(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, login)
}

// 测试用uri：localhost:8081/auth/loginUri/admin/888888
func (a *AuthController) LoginUri(ctx *gin.Context) {
	var login Login
	// 按照JSON格式绑定到对象
	if err := ctx.ShouldBindUri(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, login)
}
