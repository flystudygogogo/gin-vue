package chapter02

import (
	"fmt"
	"github.com/gin-gonic/gin"
	__ "go_code/gin-vue/proto"
	"net/http"
)

func OutAsciiJson(ctx *gin.Context) {
	ctx.AsciiJSON(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>hello,world!</b>",
	})
}

func OutXml(ctx *gin.Context) {
	ctx.XML(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>hello,world!</b>",
	})
}

func OutYaml(ctx *gin.Context) {
	ctx.YAML(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>hello,world!</b>",
	})
}

func OutProto(ctx *gin.Context) {
	userData := &__.User{Name: "tom", Age: 12}
	ctx.ProtoBuf(http.StatusOK, userData)
}

//路由跳转

func RedirectA(ctx *gin.Context) {
	fmt.Println("这是A路由")
	//状态码必须为302
	//ctx.Redirect(http.StatusFound, "/redirect_b")
	ctx.Redirect(http.StatusFound, "https://www.baidu.com/")
}

func RedirectB(ctx *gin.Context) {
	fmt.Println("这是B路由")
	ctx.String(http.StatusOK, "这是B路由")
}
