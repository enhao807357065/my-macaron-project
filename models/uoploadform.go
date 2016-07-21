package models

import "mime/multipart"

type UploadForm struct {
	Title      	string                	`form:"title"`
	TextUpload 	*multipart.FileHeader 	`form:"file"`
}