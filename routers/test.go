package routers

import (
	"gopkg.in/macaron.v1"
	"net/http"
	"fmt"
	"strings"
)

func Test(ctx *macaron.Context) {
	ctx.Data["Title"] = "我爱北京天安门"
	ctx.HTML(http.StatusOK, "test")
}

func Test1(ctx *macaron.Context) {
	//i, err := strconv.Atoi("")
	//fmt.Println("iii", i)
	//if err != nil {
	//	log.Printf("%v", err.Error())
	//}

	ctx.Data["Title"] = "我爱北京天安门"
	ctx.HTML(http.StatusOK, "test")
}

func UnicodeIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	fmt.Println("result: ", result)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		fmt.Println("prefix: ", prefix)
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		fmt.Println("rs: ", rs)
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}

	return result
}