package chapter05

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func MiddleWare1(ctx *gin.Context) {

	startTime := time.Now()

	fmt.Println("这是自定义中间件1--开始")
	ctx.Next()
	timeUse := time.Since(startTime)
	fmt.Println(timeUse)

	fmt.Println("这是自定义中间件1--结束")
}

func MiddleWare2() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("这是自定义中间件2--开始")
		////context.Next()
		//if 3 < 4 {
		//	context.Abort()
		//}
		time.Sleep(3)
		context.Next()

		fmt.Println("这是自定义中间件2--结束")
	}
}

func MiddleWare3(ctx *gin.Context) {
	fmt.Println("这是自定义中间件3--开始")
	time.Sleep(3)
	ctx.Next()
	fmt.Println("这是自定义中间件3--结束")
}
