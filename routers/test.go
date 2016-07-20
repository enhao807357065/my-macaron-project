package routers

import (
	"gopkg.in/macaron.v1"
	"net/http"
	"strings"
	"adwall/services"
	"strconv"
	"fmt"
	"os"
	"io"
	"log"
	"path"
	"time"
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

func UploadFile(ctx *macaron.Context) {
	file, fh, err := ctx.GetFile("file")
	shortPath := ""
	if file != nil && err == nil {
		defer file.Close()
		ext := path.Ext(fh.Filename)
		name := strconv.FormatInt(time.Now().UnixNano(), 36)

		dateDir := time.Now().Format("2006_01_02") + "/"
		if _, err := os.Stat("upload/" + dateDir); os.IsNotExist(err) { //如果文件夹不存在则创建
			os.MkdirAll("upload/" + dateDir, 0766)
		}
		shortPath = dateDir + name + ext
		path := "upload/" + shortPath

		write, oe := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0766)
		if oe != nil {
			log.Printf("%v", oe.Error())
			panic(oe)
		}

		err := PipeAndClose(file, write)

		if err != nil {
			if err != io.EOF {
				ctx.Status(500)
				return
			}
		}
	}
	result := map[string]string{
		"path": shortPath,
	}
	ctx.JSON(http.StatusOK, result)
}

func PipeAndClose(in io.ReadCloser, out io.WriteCloser) error {
	defer func() {
		in.Close()
		out.Close()
	}()
	return Pipe(in, out)
}

func Pipe(in io.Reader, out io.Writer) error {
	buf := make([]byte, 2048)
	for {
		n, err := in.Read(buf)

		if n > 0 {
			_, err = out.Write(buf)
		}
		if err != nil {
			if err != io.EOF {
				return err
			}
			return nil
		}
	}
	return nil
}