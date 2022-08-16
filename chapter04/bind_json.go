package chapter04

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ToBindJson(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter04/bind_json.html", nil)
}

func DoBindJson(ctx *gin.Context) {
	var user User
	err := ctx.ShouldBind(&user)
	fmt.Println(err)

	fmt.Println(user)
	
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusNotFound,
			"msg":  "失败",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "成功",
		})
	}
}
