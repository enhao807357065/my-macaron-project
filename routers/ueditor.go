package routers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"gopkg.in/macaron.v1"
	"adwall/models"
	"adwall/util"
	"log"
)

func GetTextHandler(ctx *macaron.Context, w http.ResponseWriter, r *http.Request) {
	action := r.URL.Query()["action"][0]
	if r.Method == "GET" {
		if action == "config" {
			config(w, r)
		}
	} else if r.Method == "POST" {
		if action == "uploadimage" {
			//			uploadImage(w, r)
		}
	}
}

func config(w http.ResponseWriter, r *http.Request) {
	imageUrl := util.GetImageUrl()
	imageActionName := util.GetImageActionName()
	imageUrlPrefix := util.GetImageUrlPrefix()
	imageFieldName := util.GetImageFieldName()
	//	imageAllowFiles := util.GetIamgeAllowFiles()
	//	fmt.Println("imageAllowFiles: ", len(imageAllowFiles[1:len(imageAllowFiles)-1]))
	imageMaxSize := util.GetImageMaxSize()

	imageAllowFiles := util.GetImageAllowFiles()

	mpstr := map[string]string{
		"imageUrl": imageUrl,
		"imageActionName": imageActionName,
		"imageUrlPrefix": imageUrlPrefix,
		"imageFieldName": imageFieldName,
		"imageAllowFiles": imageAllowFiles,
		"imageMaxSize": imageMaxSize,
	}
	x, err := json.Marshal(mpstr)
	if err != nil {
		fmt.Println(err)
		log.Printf("%v", err)
	}
	w.Write(x)
}

func PostTextHandler(ctx *macaron.Context, w http.ResponseWriter, r *http.Request, uf models.UploadForm) {
	var back string

	if uf.TextUpload != nil {
		data, err := util.UploadFile("file", uf.TextUpload, util.FullUrl(util.GetServerUrl(), "/service/upload"))
		if err != nil {
			fmt.Println("parse json err", err)
			log.Printf("%v", err)
			panic(err)
		} else {
			var obj map[string]interface{}
			err = util.ReaderToJson(data, &obj)
			if err != nil {
				log.Printf("%v", err)
				fmt.Println("parse json err", err)
			}
			if obj["path"] != nil {
				back = obj["path"].(string)
			}
		}
	}

	urlstr := util.GetImageUrlPrefix() + back
	b, err := json.Marshal(map[string]string{
		"url":      urlstr, //保存后的文件路径
		"title":    "",                                         //文件描述，对图片来说在前端会添加到title属性上
		//		"original": header.Filename,                            //原始文件名
		"state":    "SUCCESS",                                  //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
	})
	if err != nil {
		panic(err)
	}
	w.Write(b)
}
