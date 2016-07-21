package routers

import (
	"gopkg.in/macaron.v1"
	"path"
	"strconv"
	"time"
	"os"
	"net/http"
	"log"
	"io"
	"encoding/csv"
	"fmt"
	"bufio"
	"bytes"
)

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
	result := map[string]interface{}{
		"path": shortPath,
	}
	ctx.JSON(http.StatusOK, result)
}

func LocalExcel(ctx *macaron.Context) {
	prefixStr := "excel/"

	r :=  strconv.FormatInt(time.Now().UnixNano(), 36)
	dateDir := time.Now().Format("2006_01_02") + "_excel/"
	if _, err := os.Stat(prefixStr + dateDir); os.IsNotExist(err) { //如果文件夹不存在则创建
		os.MkdirAll(prefixStr + dateDir, 0766)
	}
	name := r + ".xls"

	f, err := os.Create(prefixStr + dateDir + name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f)
	w.Write([]string{"编号","姓名","年龄"})
	w.Write([]string{"1","张三","23"})
	w.Write([]string{"2","李四","24"})
	w.Write([]string{"3","王五","25"})
	w.Write([]string{"4","赵六","26"})
	w.Flush()

	fmt.Println("create success")
}

func DownloadExcel(ctx *macaron.Context, w http.ResponseWriter) {
	str := CreateStatement()
	ctx.Header().Set("Content-Type", "application/vnd.ms-excel")
	ctx.Write(str)
}

func CreateStatement() ([]byte) {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	bw.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	bw.WriteString("编号,")
	bw.WriteString("姓名,")
	bw.WriteString("年龄,")
	bw.WriteString("\n")

	bw.WriteString("1,")
	bw.WriteString("张三,")
	bw.WriteString("23,")
	bw.WriteString("\n")
	bw.WriteString("2,")
	bw.WriteString("李四,")
	bw.WriteString("24,")

	bw.WriteString("\n")
	bw.Flush()

	fmt.Println("create end------------------")
	return b.Bytes()
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