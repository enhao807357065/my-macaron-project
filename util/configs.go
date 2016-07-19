package util

import (
	"github.com/astaxie/beego/config"
	"log"
)

var appname	 	string              	//服务器外网访问地址
var httpport 		string       		//资源服务器外网访问地址
var mysqlurl		string			//mysql链接配置

func init() {
	iniconfig, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		log.Printf("%v", err.Error())
	}

	if iniconfig != nil {
		appname = iniconfig.String("appname")
		httpport = iniconfig.String("httpport")
		mysqlurl = iniconfig.String("mysqlurl")
	}
}

func GetAppName() string {
	return appname
}

func GetHttpPort() string {
	return httpport
}

func GetMysqlUrl() string {
	return mysqlurl
}