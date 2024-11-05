package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	g.Meta `path:"/file/upload" method:"post" mime:"multipart/form-data" summary:"文件上传" tags:"文件"`
	File   *ghttp.UploadFile `json:"file" type:"file" v:"required#请选择上传文件"`
}

type FileUploadRes struct {
	Name string `json:"name" dc:"文件名"`
	Url  string `json:"url" dc:"文件URL"`
}
