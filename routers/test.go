package routers

import (
	"gopkg.in/macaron.v1"
	"net/http"
	"strings"
	"adwall/services"
	"strconv"
	"fmt"
)

func Test1(ctx *macaron.Context) {
	//i, err := strconv.Atoi("")
	//fmt.Println("iii", i)
	//if err != nil {
	//	log.Printf("%v", err.Error())
	//}

	//idStr := ctx.ParamsInt64("id")
	idStr, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if err != nil {
		panic(err)
	}

	user, err := services.GetUserById(idStr)
	ctx.Data["user"] = user

	ctx.HTML(http.StatusOK, "test")
}

func Test2(ctx *macaron.Context) {
	fmt.Println("Test2------------------------")
	nameStr := ctx.Req.Request.PostFormValue("name")

	result := map[string]string{}
	if nameStr == "张飞" {
		result["res"] = "成功"
		ctx.JSON(http.StatusOK, result)
	} else {
		result["res"] = "失败"
		ctx.JSON(http.StatusOK, result)
	}
}

func UnicodeIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}
	return result
}