package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go_code/gin-vue/chapter03"
	"go_code/gin-vue/chapter04"
	"go_code/gin-vue/chapter05"
	all_router "go_code/gin-vue/router"
	"html/template"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	//router := gin.New()
	//router.Use(gin.Logger(),gin.Recovery())
	//router.Use(gin.Logger(), gin.Recovery())

	//全局中使用中间件
	router.Use(chapter05.MiddleWare1)
	router.Use(chapter05.MiddleWare2())
	router.Use(chapter05.MiddleWare3)
	router.SetFuncMap(template.FuncMap{
		"add":    chapter03.AddNum,
		"subStr": chapter03.SubStr,
		"safe":   chapter03.Safe,
	})

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation("len_valid", chapter04.LenValid)
	}

	router.LoadHTMLGlob("template/**/*")
	//router.LoadHTMLFiles("index.html", "news.html")

	router.Static("/static", "static")

	all_router.Router(router)
	//router.StaticFS("/static", http.Dir("static"))

	//v1 := router.Group("/v1")
	//{
	//	v1.GET("/", chapter01.Hello)
	//}

	//v2 := router.Group("/v2")
	//
	//{
	//	v2.GET("/user", chapter02.User)
	//}
	////各种数据类型渲染
	//router.GET("/", chapter01.Hello)
	//router.GET()
	//router.GET("/user", chapter02.User)
	//router.GET("/user_struct", chapter02.UserInfoStruct)
	//router.GET("/arr", chapter02.ArrController)
	//router.GET("/arr_struct", chapter02.ArrStruct)
	//router.GET("/map", chapter02.Map)
	//router.GET("/map_struct", chapter02.MapStruct)
	//router.GET("/slice", chapter02.SliceController)
	//
	////路径上直接加上参数值
	//router.GET("/param/:id", chapter02.Param)
	//router.GET("/param2/*id", chapter02.Param2)
	//
	////路径中使用参数名
	//router.GET("/query", chapter02.GetQueryData)
	//router.GET("query_arr", chapter02.GetQueryArr)
	//router.GET("query_map", chapter02.GetQueryMap)
	//
	//router.GET("/to_user_add", chapter02.ToUserAdd)
	//router.POST("/do_user_add", chapter02.DoUserAdd)
	//
	//router.GET("/to_user_add2", chapter02.ToUserAdd2)
	//router.POST("/do_user_add2", chapter02.DoUserAdd2)
	//
	//router.GET("/to_user_add3", chapter02.ToUserAdd3)
	//router.POST("/do_user_add3", chapter02.DoUserAdd3)
	//
	//router.GET("/to_user_add4", chapter02.ToUserAdd4)
	//router.POST("/do_user_add4", chapter02.DoUserAdd4)
	//
	//router.GET("/test_to_upload1", chapter02.ToUpload1)
	//router.POST("/test_do_upload1", chapter02.DoUpload1)
	//
	//router.GET("/test_to_upload2", chapter02.ToUpload2)
	//router.POST("/test_do_upload2", chapter02.DoUpload2)
	//
	////router.GET("/test_to_upload3", chapter02.ToUpload3)
	////router.POST("/test_do_upload3", chapter02.DoUpload3)
	//router.GET("/output", chapter02.OutAsciiJson)
	//router.GET("/outXml", chapter02.OutXml)
	//router.GET("/outYaml", chapter02.OutYaml)
	//router.GET("/outProto", chapter02.OutProto)
	//
	//router.GET("/redirect_a", chapter02.RedirectA)
	//router.GET("/redirect_b", chapter02.RedirectB)

	//模板语法
	//router.GET("/test", chapter03.TestSyntaxTpl)

	//数据绑定
	//router.GET("/to_bind_form", chapter04.ToBindForm)
	//router.POST("/do_bind_form", chapter04.DoBindForm)
	//
	//router.GET("/to_bind_query", chapter04.BindQueryString)
	//
	//router.GET("/to_bind_json", chapter04.ToBindJson)
	//router.POST("/do_bind_json", chapter04.DoBindJson)
	//
	//router.GET("/bind_uri/:name/:age/:addr", chapter04.BindUri)
	//
	////数据校验
	//router.GET("/to_valid", chapter04.ToValidData)
	//router.POST("/do_valid", chapter04.DoValidData)

	//http.ListenAndServe(":8899", router)
	//router.Run(":8899")

	//自定义http配置
	s := http.Server{
		Addr:         ":8899",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	err := s.ListenAndServe()
	if err != nil {
		return
	}

}
