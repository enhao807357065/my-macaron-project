package main

import (
	_ "adwall/routers"
	"gopkg.in/macaron.v1"
	"github.com/go-macaron/gzip"
	"html/template"
	"adwall/routers"
	"runtime"
	"flag"
	"os"
	"fmt"
	"log"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	m := macaron.Classic()
	m.Use(macaron.Recovery())
	m.Use(macaron.Logger())
	m.Use(gzip.Gziper())
	m.Use(macaron.Static("static"))

	m.Use(macaron.Renderer(macaron.RenderOptions{
		Funcs:      []template.FuncMap{map[string]interface{}{
			"URLFor": m.URLFor,
		}},
		Directory:		"views",
		Extensions: 		[]string{".tpl", ".html"}, // Specify extensions to load for templates.
		Charset:         	"UTF-8",     // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON:      	true,        // Output human readable JSON
		IndentXML:       	true,        // Output human readable XML
		HTMLContentType: 	"text/html",
		Delims: 		macaron.Delims{"{{", "}}"}, // 模板语法分隔符，默认为 ["{{", "}}"]
	}))

	routers.RegistAllRouters(m)

	//beego.SetLogger("file", `{"filename":"project.log","level":1,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)

	//log := logs.NewLogger(logs.LevelAlert)
	//log.SetLogger("file", `{"filename":"project.log","level":1,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	//log.Async()
	//log.EnableFuncCallDepth(true)

	//set logfile Stdout
	logFile, logErr := os.OpenFile("project.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		os.Exit(1)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	m.Run(8088)
}