package chapter02

import "github.com/gin-gonic/gin"

func Router(chap02 *gin.RouterGroup) {
	chap02.GET("/user", User)
	chap02.GET("/user_struct", UserInfoStruct)
	chap02.GET("/arr", ArrController)
	chap02.GET("/arr_struct", ArrStruct)
	chap02.GET("/map", Map)
	chap02.GET("/map_struct", MapStruct)
	chap02.GET("/slice", SliceController)

	//路径上直接加上参数值
	chap02.GET("/param/:id", Param)
	chap02.GET("/param2/*id", Param2)

	//路径中使用参数名
	chap02.GET("/query", GetQueryData)
	chap02.GET("query_arr", GetQueryArr)
	chap02.GET("query_map", GetQueryMap)

	chap02.GET("/to_user_add", ToUserAdd)
	chap02.POST("/do_user_add", DoUserAdd)

	chap02.GET("/to_user_add2", ToUserAdd2)
	chap02.POST("/do_user_add2", DoUserAdd2)

	chap02.GET("/to_user_add3", ToUserAdd3)
	chap02.POST("/do_user_add3", DoUserAdd3)

	chap02.GET("/to_user_add4", ToUserAdd4)
	chap02.POST("/do_user_add4", DoUserAdd4)

	chap02.GET("/test_to_upload1", ToUpload1)
	chap02.POST("/test_do_upload1", DoUpload1)

	chap02.GET("/test_to_upload2", ToUpload2)
	chap02.POST("/test_do_upload2", DoUpload2)

	//router.GET("/test_to_upload3", chapter02.ToUpload3)
	//router.POST("/test_do_upload3", chapter02.DoUpload3)
	chap02.GET("/output", OutAsciiJson)
	chap02.GET("/outXml", OutXml)
	chap02.GET("/outYaml", OutYaml)
	chap02.GET("/outProto", OutProto)

	chap02.GET("/redirect_a", RedirectA)
	chap02.GET("/redirect_b", RedirectB)

}
