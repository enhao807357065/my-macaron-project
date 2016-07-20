package routers

import (
	"gopkg.in/macaron.v1"
	"net/http"
)

func RegistAllRouters(ctx *macaron.Macaron) {

	ctx.NotFound(func (ctx *macaron.Context) {
		ctx.Status(http.StatusNotFound)
	})

	ctx.Get("/", func(ctx *macaron.Context) {
		ctx.Write([]byte("server is running ok!"))
	})

	ctx.Get("/test1", Test1)
	ctx.Post("/post/form", Test2)
	ctx.Post("/file_upload", UploadFile)
}
