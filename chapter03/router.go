package chapter03

import "github.com/gin-gonic/gin"

func Router(chap03 *gin.RouterGroup) {
	chap03.GET("/test", TestSyntaxTpl)
}
