package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadImgToCloudReq struct {
	g.Meta `path:"/upload/cloud" method:"post" summary:"上传图片到云存储" mime:"multipart/form-data" tags:"文件"`
	File   *ghttp.UploadFile `json:"file" type:"file" v:"required#请选择上传文件"`
}

type UploadImgToCloudRes struct {
	Url string `json:"url" dc:"文件URL"`
}
