package routers

import (
	"gopkg.in/macaron.v1"
	"net/http"
)

func RegistAllRouters(ctx *macaron.Macaron) {

	ctx.NotFound(func(ctx *macaron.Context) {
		ctx.Status(http.StatusNotFound)
	})

	ctx.Get("/", func(ctx *macaron.Context) {
		ctx.Write([]byte("server is running ok!"))
	})

	ctx.Get("/test1", Test1)
	ctx.Post("/post/form", Test2)

	//一般服务
	ctx.Group("/service", func() {
		//文件上传
		ctx.Post("/upload", UploadFile)
		//本地生成excel
		ctx.Get("/local/excel", LocalExcel)
		//流的形式生成excel下载  不生成excel文件
		ctx.Get("/download/excel", DownloadExcel)
	})
}
