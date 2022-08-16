package chapter04

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BindUri(ctx *gin.Context) {
	var user User

	err := ctx.ShouldBindUri(&user)
	fmt.Println(user)

	if err != nil {
		ctx.String(http.StatusNotFound, "绑定失败")
	} else {
		ctx.String(http.StatusOK, "绑定成功")
	}
}
