package model

import "github.com/gogf/gf/v2/net/ghttp"

type FileUploadInput struct {
	File       *ghttp.UploadFile `json:"file"`
	Name       string            `json:"name"`
	RandomName bool              `json:"random_name"`
}

type FileUploadOutput struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Src  string `json:"src"`
}
