package chapter02

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Id   int    `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
	Age  int    `form:"age" json:"age"`
	Addr string `form:"addr" json:"addr"`
}

func User(ctx *gin.Context) {

	name := "fallen"
	ctx.HTML(200, "User/user.html", name)
}

//结构体渲染

func UserInfoStruct(ctx *gin.Context) {
	user_info := UserInfo{Id: 1, Name: "fallen", Age: 19, Addr: "xxx"}
	ctx.HTML(http.StatusOK, "chapter02/user_info.html", user_info)
}

//数组渲染

func ArrController(ctx *gin.Context) {
	arr := [3]int{1, 3, 4}
	ctx.HTML(http.StatusOK, "chapter02/arr.html", arr)
}

//结构体数组渲染

func ArrStruct(ctx *gin.Context) {

	arrStruct := [3]UserInfo{
		{Id: 1, Name: "tom", Age: 18, Addr: "xxx"},
		{Id: 2, Name: "tom2", Age: 19, Addr: "xxx2"},
		{Id: 1, Name: "tom3", Age: 20, Addr: "xxx3"},
	}

	ctx.HTML(http.StatusOK, "chapter02/arr_struct.html", arrStruct)
}

//map渲染

func Map(ctx *gin.Context) {

	mapData := map[string]string{
		"age":  "18",
		"name": "tom",
	}

	mapData2 := map[string]interface{}{"mapData": mapData}
	ctx.HTML(http.StatusOK, "chapter02/map.html", mapData2)
}

//map结构体渲染

func MapStruct(ctx *gin.Context) {
	mapStruct := map[string]UserInfo{"user": {Id: 1, Name: "tom", Age: 18, Addr: "xxx"}}
	ctx.HTML(http.StatusOK, "chapter02/map_struct.html", mapStruct)
}

//切片渲染

func SliceController(ctx *gin.Context) {
	sliceData := []int{1, 2, 3, 4, 5}
	ctx.HTML(http.StatusOK, "chapter02/slice.html", sliceData)
}

//路径上直接加上参数值

func Param(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.String(http.StatusOK, "hello,%s", id)
}

func Param2(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.String(http.StatusOK, "hello,%s", id)
}

//路径中使用参数名

func GetQueryData(ctx *gin.Context) {
	id := ctx.Query("id")

	name := ctx.DefaultQuery("name", "tom")
	fmt.Println(name)
	ctx.String(http.StatusOK, "tom,%s", id)
}

func GetQueryArr(ctx *gin.Context) {
	ids := ctx.QueryArray("ids")

	ctx.String(http.StatusOK, "tom,%s", ids)
}

func GetQueryMap(ctx *gin.Context) {
	user := ctx.QueryMap("user")

	ctx.String(http.StatusOK, "tom,%s", user)
}

//post请求
//到用户添加页面

func ToUserAdd(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter02/user_add.html", nil)
}

//添加用户

func DoUserAdd(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Println(username)
	fmt.Println(password)

	ctx.String(http.StatusOK, "添加成功")
}

func ToUserAdd2(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter02/user_add2.html", nil)
}

func DoUserAdd2(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	age := ctx.DefaultPostForm("age", "18")

	loves := ctx.PostFormArray("love")
	user := ctx.PostFormMap("user")
	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(age)
	fmt.Println(loves)
	fmt.Println(user)

	ctx.String(http.StatusOK, "添加成功")
}

func ToUserAdd3(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter02/user_add3.html", nil)
}

func DoUserAdd3(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	age := ctx.PostForm("age")

	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(age)
	//map_data := map[string]interface{}{
	//	"code": 200,
	//	"msg":  "成功",
	//}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功"})
}

func ToUserAdd4(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter02/user_add4.html", nil)
}

func DoUserAdd4(ctx *gin.Context) {
	var userInfo UserInfo
	err := ctx.ShouldBind(&userInfo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userInfo)

	ctx.String(http.StatusOK, "绑定成功")
}
