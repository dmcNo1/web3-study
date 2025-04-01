package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetMsg(ctx *gin.Context) {
	time.Sleep(time.Second * 3)
	ctx.String(http.StatusOK, "结束了，看日志吧")
}
