package chapter03

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

type Article struct {
	Id    int
	Title string
	Desc  string
}

func TestSyntaxTpl(ctx *gin.Context) {
	name := "tom"
	arr := []int{11, 22, 33, 44}

	timeNow := time.Now()
	fmt.Println(timeNow)
	article := Article{Id: 1, Title: "西游记", Desc: "西游记外传，不容错过"}
	map_data := map[string]interface{}{
		"name":    name,
		"arr":     arr,
		"birth":   "2002-01-16",
		"article": article,
		"nowTime": timeNow,
	}
	ctx.HTML(http.StatusOK, "chapter03/test.html", map_data)
}

//定义函数

func AddNum(num1 int, num2 int) int {
	return num1 + num2
}

//自定义模板函数实战

func SubStr(str string, num int) string {
	strLen := len(str)
	//如果字符串长度小于等于指定长度，直接返回字符串
	if strLen <= num {
		return str
	} else {
		newStr := str[0:num]
		return newStr + "..."
	}

}

func Safe(str string) template.HTML {

	return template.HTML(str)
}
