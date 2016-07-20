package util

import (
	"github.com/astaxie/beego/config"
	"log"
	"strings"
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