package util

import (
	"github.com/astaxie/beego/config"
	"log"
	"strings"
	"mime/multipart"
	"fmt"
	"net/http"
	"bytes"
	"io"
	"io/ioutil"
	"encoding/json"
)

var APPNAME	 	string              	//服务器外网访问地址
var HTTPPPORT 		string       		//资源服务器外网访问地址
var MYSQLURL		string			//mysql链接配置
var SERVERURL		string			//服务器地址

var IMAGEURL		string			//富文本编辑器上传图片
var IMAGEACTIONNAME	string			//上传图片action
var IMAGEURLPREFIX	string			//图片上传成功返回给前端的前缀
var IMAGEFIELDNAME	string			//图片input的name
var IMAGEMAXSIZE	string			//图片大小  单位b
var IMAGEALLOWFILES	string			//图片格式

func init() {
	iniconfig, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		log.Printf("%v", err.Error())
	}

	if iniconfig != nil {
		APPNAME = iniconfig.String("appname")
		HTTPPPORT = iniconfig.String("httpport")
		MYSQLURL = iniconfig.String("mysqlurl")
		IMAGEURL = iniconfig.String("imageurl")
		IMAGEACTIONNAME = iniconfig.String("imageactionname")
		IMAGEURLPREFIX = iniconfig.String("imageurlprefix")
		IMAGEFIELDNAME = iniconfig.String("imagefiledname")
		IMAGEMAXSIZE = iniconfig.String("imagemaxsize")
		IMAGEALLOWFILES = iniconfig.String("imageallowfiles")
		SERVERURL = iniconfig.String("serverurl")
	}
}

func GetAppName() string {
	return APPNAME
}

func GetHttpPort() string {
	return HTTPPPORT
}

func GetMysqlUrl() string {
	return MYSQLURL
}

func GetImageUrl() string {
	return IMAGEURL
}

func GetImageActionName() string {
	return IMAGEACTIONNAME
}

func GetImageUrlPrefix() string {
	return IMAGEURLPREFIX
}

func GetImageFieldName() string {
	return IMAGEFIELDNAME
}

func GetImageMaxSize() string {
	return IMAGEMAXSIZE
}

func GetImageAllowFiles() string {
	return IMAGEALLOWFILES
}

func GetServerUrl() string {
	return SERVERURL
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

func UploadFile(field string, header *multipart.FileHeader, url string) (body io.ReadCloser, err error) {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	fw, err := w.CreateFormFile(field, header.Filename)

	if err != nil {
		fmt.Println("c")
		return nil, err
	}

	file, e := header.Open()
	if e != nil {
		return nil, e
	}
	defer file.Close()
	_, err = io.Copy(fw, file)
	if err != nil {
		fmt.Println("e")
		return nil, err
	}

	w.Close()
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		fmt.Println("f")
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("g")
		return nil, err
	}
	return res.Body, err
}

func ReaderToJson(reader io.Reader, v interface{}) error {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)
	return err
}

func ToJsonString(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}

func FullUrl(server string, path string) string {
	if path == ""{
		return ""
	}
	if len(path) > 9 && (path[:7] == "http://" || path[:8] == "https://") {
		return path
	}
	if path[0] == '/' {
		return server + path[1:]
	}
	return server + path
}