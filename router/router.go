package all_router

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue/chapter01"
	"go_code/gin-vue/chapter02"
	"go_code/gin-vue/chapter03"
	"go_code/gin-vue/chapter04"
	"go_code/gin-vue/chapter05"
)

func Router(router *gin.Engine) {
	chap01 := router.Group("/chapter01")
	chap02 := router.Group("/chapter02")
	chap03 := router.Group("/chapter03")
	chap04 := router.Group("/chapter04")
	chap05 := router.Group("/chapter05")

	chapter01.Router(chap01)
	chapter02.Router(chap02)
	chapter03.Router(chap03)
	chapter04.Router(chap04)
	chapter05.Router(chap05)

}
