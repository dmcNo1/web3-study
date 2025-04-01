package user

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

// 获取参数，localhost:8081/user/zhagnsan/coding?level=slowly
func (u *UserController) Action(ctx *gin.Context) {
	// Restful风格可以通过func (c *Context) Param(key string) string获取
	name := ctx.Param("name")
	action := strings.Trim(ctx.Param("action"), "/")
	// 获取Get中的传参，可以用Query()或者DefaultQuery()
	level := ctx.DefaultQuery("level", "fast")
	ctx.String(http.StatusOK, fmt.Sprintf("%s is %s %s", name, action, level))
}

// 获取表单中的参数
func (u *UserController) Form(ctx *gin.Context) {
	// DefaultPostForm()或者PostForm()获取表单中的数据
	method := ctx.DefaultPostForm("method", "post")
	username := ctx.PostForm("username")
	password := ctx.PostForm("userpassword")
	ctx.String(http.StatusOK, fmt.Sprintf("method:%s --- username:%s password:%s", method, username, password))
}

// 上传单个文件
func (u *UserController) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(http.StatusInternalServerError, "上传文件异常")
	}
	// 获取文件名
	destPath := filepath.Join("./files/", file.Filename)
	// 保存文件
	err = ctx.SaveUploadedFile(file, destPath)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "文件保存异常")
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success!"})
}

// 上传多个文件
func (u *UserController) UploadFiles(ctx *gin.Context) {
	// 获取MultipartForm，然后从中获取到上传的文件列表
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "上传文件异常")
	}
	files := form.File["files"]
	// 遍历文件列表，逐个保存
	for _, file := range files {
		destPath := filepath.Join("./files/", file.Filename)
		err := ctx.SaveUploadedFile(file, destPath)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "文件保存异常")
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success!"})
}
