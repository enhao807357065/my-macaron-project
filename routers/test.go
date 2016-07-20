package routers

import (
	"gopkg.in/macaron.v1"
	"net/http"
	"adwall/services"
	"strconv"
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
