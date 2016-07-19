package main

import (
	"adwall/routers"
	_ "adwall/routers"
	_ "adwall/util"
	"flag"
	"fmt"
	"github.com/go-macaron/gzip"
	"gopkg.in/macaron.v1"
	"html/template"
	"log"
	"os"
	"runtime"
	"adwall/util"
	"strconv"
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
		Funcs: []template.FuncMap{map[string]interface{}{
			"URLFor": m.URLFor,
		}},
		Directory:       "views",
		Extensions:      []string{".tpl", ".html"}, // Specify extensions to load for templates.
		Charset:         "UTF-8",                   // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON:      true,                      // Output human readable JSON
		IndentXML:       true,                      // Output human readable XML
		HTMLContentType: "text/html",
		Delims:          macaron.Delims{"{{", "}}"}, // 模板语法分隔符，默认为 ["{{", "}}"]
	}))

	routers.RegistAllRouters(m)

	//set logfile Stdout
	logFile, logErr := os.OpenFile("project.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		os.Exit(1)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	httpPort, _ := strconv.Atoi(util.GetHttpPort())
	m.Run(httpPort)
}
