package chapter04

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

//binding

//type Article struct {
//	Id      int     `form:"-"`
//	Title   string  `form:"title" binding:"len_valid"`
//	Content string  `form:"content"`
//	Desc    string  `form:"desc"`
//	Uid     []User  `binding:"required,dive,min=10"`
//	Name    [][]int `binding:"min=10,dive,max=20,dive,required""`
//}

//valid for beego

type Article struct {
	Id      int    `form:"-"`
	Title   string `form:"title" valid:"Required"`
	Content string `form:"content"`
	Desc    string `form:"desc"`
}

func ToValidData(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter04/valid_data.html", nil)
}

//binging

//func DoValidData(ctx *gin.Context) {
//	var article Article
//
//	err := ctx.ShouldBind(&article)
//	fmt.Println(err)
//	fmt.Println(article)
//	ctx.String(http.StatusOK, "成功")
//
//}

//beego

func DoValidData(ctx *gin.Context) {
	var article Article

	err := ctx.ShouldBind(&article)
	fmt.Println(err)
	fmt.Println(article)

	//初始化验证机器
	valid := validation.Validation{}
	b, err1 := valid.Valid(&article)
	fmt.Println(err1)
	if !b {
		//校验有错误
		for _, err2 := range valid.Errors {
			fmt.Println(err2.Key)
			fmt.Println(err2.Message)
		}
	}
	ctx.String(http.StatusOK, "成功")

}

//判断是不是大于6

var LenValid validator.Func = func(fl validator.FieldLevel) bool {
	data := fl.Field().Interface().(string)
	fmt.Println(data)

	if len(data) > 6 {
		return true
	}
	return false
}
